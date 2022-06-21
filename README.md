# Illuminate contest details
- $50,000 USDC main award pot
- $5,000 USDC gas optimization award pot
- Join [C4 Discord](https://discord.gg/EY5dvm3evD) to register
- Submit findings [using the C4 form](https://code4rena.com/contests/2022-06-illuminate-contest/submit) -- Update link
- [Read our guidelines for more details](https://docs.code4rena.com/roles/wardens)
- Starts June 21, 2022 20:00 UTC
- Ends June 26, 2022 19:59 UTC

| **Contracts**    | **Link** | **LOC** | **LIBS** | **External** |
|--------------|------|------|------|------|
| Marketplace |[Link](https://github.com/code-423n4/2022-06-illuminate/blob/main/marketplace/MarketPlace.sol)| 253 | [Safe.sol](https://github.com/code-423n4/2022-06-illuminate/blob/main/marketplace/Safe.sol) | [External.md](https://github.com/code-423n4/2022-06-illuminate/blob/main/external.md) |
| Lender |[Link](https://github.com/code-423n4/2022-06-illuminate/blob/main/lender/Lender.sol)| 738 | [Safe.sol](https://github.com/code-423n4/2022-06-illuminate/blob/main/lender/Safe.sol) | [External.md](https://github.com/code-423n4/2022-06-illuminate/blob/main/external.md)
| Redeemer |[Link](https://github.com/code-423n4/2022-06-illuminate/blob/main/redeemer/Redeemer.sol)| 306 | [Safe.sol](https://github.com/code-423n4/2022-06-illuminate/blob/main/redeemer/Safe.sol) | [External.md](https://github.com/code-423n4/2022-06-illuminate/blob/main/external.md) |


# Introduction
Illuminate is a fixed-rate lending protocol designed to aggregate other fixed-rate protocols' guaranteed yields (in the form of principal tokens), providing Illuminate's users a loose guarantee of the best rate, and deepening liquidity across the fixed-rate space.

Most simply described, Illuminate aggregates and wraps principal tokens with similar maturities and underlying assets into one single (meta) principal token (iPT).

A front-end interface provides an aggregated _best rate_ by purchasing and wrapping external tokens.

While the wrapped / meta principal token (iPT) is traded on a secondary market through a YieldSpace AMM (out of scope for this audit) providing an on-chain guarantee of the _best rate_. 

As arbitrageurs can easily purchase external PTs -> mint & sell iPTs, variance in external principal token prices is reduced, while a secondary market for iPTs can be confident it provides the best rate and depth.

## Main Participant Interactions

**On-Chain lenders:** Purchase iPTs through Marketplace.sol which serves as a router for a secondary market's YieldSpace AMM.

**Off-Chain lenders:** `lend` on Lender.sol, with the assumption an off-chain interface can identify which principal token to directly purchase and then input the correct parameters.

**Arbitrageurs:** Purchase external principal tokens, `mint` iPTs on Lender.sol, and then sell iPTs through the Marketplace.sol which acts as a router.

**Redemption:** 

iPT redemption itself is straightforward as users redeem 1 iPT for 1 underlying token upon maturity.

This first requires the redemption of any externally wrapped PTs that Lender.sol may hold. These external PT `redeem` methods are public, however it is expected that the admins will likely execute these leaving the users solely responsible for redeeming their iPTs.

iPTs can be redeemed either with a `redeem` method on Redeemer.sol, or through the authorized iPT [EIP-5095](https://github.com/ethereum/EIPs/pull/5095) `redeem` method.

**Admins:**
- Set up contracts
- Create markets
- Retain emergency withdraw privileges (72 hr delay)
- Set fees

-----------------------------------

## Links:

General Project Docs: https://docs.illuminate.finance

Contract Docs: https://docs.illuminate.finance/smart-contracts

-----------------------------------

## Important Note:

When it comes to input sanitization, we err on the side of having external interfaces validate their input, rather than socializing costs to do checks such as:
- Checking for address(0)
- Checking for input amounts of 0
- Checking to ensure the protocol enum matches the parameters provided
- Any similar input sanitization

-----------------------------------

## Other Information / Interaction Paths 

#### Setting Up Markets
A market, defined by an underlying asset (`address`) and a maturity (`uint256`) is created by the admin via the MarketPlace contract. 

After deploying a market there are a number of necessary approvals:
- Lender.sol approves all external protocols to take underlying.
- Lender.sol approves Aave to take underlying (for APWine's deposits)
- Lender.sol approves Redeemer.sol to take all external principal tokens.

#### Redemption Specifics
As external principal tokens mature, the admin (or others) `redeem` the principal tokens using `Redeemer.sol`.

`redeem` methods transfer held PTs from `Lender.sol` to `Redeemer.sol` and then operate the appropriate external redemption methods. Should the iPT mature _after_ all external PTs mature, each iPT will then have equivalent backing of underlying tokens.

As the iPT then matures, lenders can redeem their capital with a `redeem` method on Redeemer.sol, or through the authorized iPT [EIP-5095](https://github.com/ethereum/EIPs/pull/5095) `redeem` method.

-----------------------------------

# Contract Descriptions / Areas of Concern

## **Lender:**

`Lender.sol` facilitates all lending and minting via an overloaded `lend` method and generic `mint` method. Off-chain interfaces identify an external principal token to purchase/wrap and then provide corresponding parameters to the `lend` method. 

When `lend` is called, underlying assets are transferred to `Lender.sol`. From there, `Lender.sol` purchases an external protocol's principal tokens and mints the user (msg.sender) an equivalent amount of wrapped/meta principal tokens (iPTs).

The admin is given privilages to pause specific external principal token wrapping. This may be important in the event of external protocol insolvency or bugs. In addition, the admin is able to able to withdraw fees.

### Areas of Concern

- The admin must not be able to withdraw more fees than what he is entitled to and fee calculation is correct
- There are no ways to abuse the parameters for external contracts or pools, mismatching and allowing the wrapping of either:
        - An external PT with a maturity > that of the iPT
        - An external PT with a different underlying than that of the iPT
- There are no adverse effects from calling the wrong principal for a given `lend` function to other users.
- The approval process is sound. Once `createMarket` is executed, users are not able to mint meta-principal tokens until `approve` has approved the Lender for all principal tokens in that market.
-----------------------------------

## **Marketplace:**

`Marketplace.sol` acts as the central hub for creating new `markets` (defined as an asset-matury pair), and the initial contract creation for iPTs which conform to the ERC5095 interface. The `markets` mapping is stored and admins are the only ones that can create markets.

In addition, `MarketPlace.sol` acts as a "Router" contract for our YieldSpace markets and the secondary market purchase/sale of iPTs. The router simply tracks live pools alongside their underlying <-> maturity via the `pools` mapping that can only be changed by the admin.

### Areas of Concern

- Users are able to swap between their meta-principal tokens and the underlying using YieldSpace pools.
- There are no oversights in the creation of new markets

-----------------------------------

## **Redeemer:**

Redemption can generally be split into two categories, the redemption of external principal tokens currently held on `Lender.sol`, and then the redemption of underlying tokens to those holding currently matured iPTs. Both of these redemptions are covered under various overloaded `redeem` methods.

As external principal tokens mature, anyone can call the corresponding public `redeem` methods. This transfers external principal tokens from `Lender.sol` to `Redeemer.sol` at which point external principal tokens are redeemed with their corresponding protocol unique methods, and underlying tokens then sit in `Redeemer.sol`.

As iPTs mature, users can can call the corresponding public `redeem` method with an enum of 0 to then burn their iPT balance and claim a 1-1 amount of the underlying that sits in `Redeemer.sol`.

In order to comply with ERC-5095, we have also included an `authRedeem` method which can be called solely by the iPT itself, similarly burning and claiming underlying.

### Areas of Concern

- Ensure that something out-of-order would not enable a user to redeem more or less tokens than they are entitled to.
- The `authRedeem` methods could not be used to incorrectly redeem funds to another user or steal funds in some manner.
-----------------------------------
