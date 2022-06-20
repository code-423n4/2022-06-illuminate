# Illuminate contest details
TODO: Update this section
- $67,500 worth of ETH main award pot
- $7,500 worth of ETH gas optimization award pot
- Join [C4 Discord](https://discord.gg/EY5dvm3evD) to register
- Submit findings [using the C4 form](https://code423n4.com/2021-09-swivel-contest/submit)
- [Read our guidelines for more details](https://docs.code4rena.com/roles/wardens)
- Starts September 30, 2021 00:00 UTC
- Ends October 6, 2021 23:59 UTC

# Introduction
TODO: Julian will write this out

General Project Docs:https://docs.illuminate.finance

Contract Docs: https://docs.illuminate.finance/developers/contract 


### **Lend and Redeem Path:**

A market, defined by an underlying asset (`address`) and a maturity (`uint256`) is created by the admin via the MarketPlace contract. From there, the admin approves the Lender for all principal tokens via the Lender contract.

Next, a user can lend capital for that market by calling `lend`. Depending on the principal he passes to `lend`, he will have to provide different parameters. When `lend` is called, his underlying asset is transferred to the Lender. From there, the Lender converts the underlying asset into principal tokens for the chosen principal. Finally, the `lend` operation ends by minting meta-principal tokens to the user. These tokens will represent his loan in this particular market.

Once the loan has matured, the admin will call `redeem` on all principals servicing a given market. These `redeem` methods will receive the owed principal tokens from the Lender contract and call the redemption functions for each principal. At that point, they will be holding the originally lent capital plus interest.

Next, a user can redeem his owed capital by calling `redeem` with Illuminate's principal. This will result in his meta-principal tokens being burned, and a proportional amount of underlying assets being returned to him.

### ** metaPrincipal token [ERC5095](https://github.com/ethereum/EIPs/pull/5095) functionality:**

The meta-principal token represents a user's owed capital in a given market, defined by an underlying asset and maturity. Meta-principal tokens are minted to users when they call `lend`. The number of tokens they own in a given market indicates how much underlying assets they are entitled to when the market matures.

A meta-principal token is minted for each market.

The meta-principal token conforms to the ERC5095 interface. The purpose of this interface is to indicate ownership of an underlying ERC20 at a future date. In the case of Illuminate, this is used to represent capital that has been lent. 

By having a common token accross all principals for a given market, Illuminate enables arbitrage across rates, ensuring that users get access to the best possible rate. In addition, having a common token for a market allows it to be traded for its underlying via YieldSpace pools.

# Smart Contracts 

## External Contracts

Below are the external contracts used by Lender.sol, Redeemer.sol and MarketPlace.sol.

### Lender

- [`ISwivel(swivelAddr).initiate(o, a, s);`](ISwivel(swivelAddr).initiate(o, a, s);)
- [`IElementToken(principal).underlying() != u`](https://github.com/element-fi/elf-contracts/blob/885666433894c598223ea6e32f8cf38236efc2f1/contracts/Tranche.sol#L28)
- [`IElementToken(principal).unlockTimestamp()`](https://github.com/element-fi/elf-contracts/blob/885666433894c598223ea6e32f8cf38236efc2f1/contracts/Tranche.sol#L18)
- [`uint256 purchased = IElement(e).swap(swap, fund, r, d);`](https://github.com/element-fi/elf-contracts/blob/65fddc8e750e156605b2d7e01ceffd4bbcb8c978/contracts/interfaces/IVault.sol#L35)
- [`token.yieldToken() != u`](https://github.com/pendle-finance/pendle-core/blob/b34d265e4fe8e3a6f79bdec1ab88ab2fd49a882c/contracts/interfaces/IPendleYieldTokenHolder.sol#L30) (Pendle)
- [`token.expiry() > m`](https://github.com/pendle-finance/pendle-core/blob/b34d265e4fe8e3a6f79bdec1ab88ab2fd49a882c/contracts/interfaces/IPendleYieldTokenHolder.sol#L36) (Pendle)
- [`uint256 returned = IPendle(pendleAddr).swapExactTokensForTokens(a - calculateFee(a), r, path, address(this), d)[0];`](https://github.com/pendle-finance/pendle-core/blob/01154885472188022518745f2ca08104955b8b2b/contracts/interfaces/IUniswapV2Router01.sol#L127)
- [`ITempus(principal).yieldBearingToken() != IErc20Metadata(u)`](https://github.com/tempus-finance/tempus-protocol/blob/f48b6f2e4563036e4292809ead1e215e6824eb16/contracts/TempusPool.sol#L32)
- [`ITempus(principal).maturityTime() > m`](https://github.com/tempus-finance/tempus-protocol/blob/f48b6f2e4563036e4292809ead1e215e6824eb16/contracts/TempusPool.sol#L36)
- [`ITempus(tempusAddr).depositAndFix(Any(x), Any(t), a - calculateFee(a), true, r, d)`](https://github.com/tempus-finance/tempus-protocol/blob/f431e821c81d8cdeae8ad433d160230563f121de/contracts/TempusController.sol#L52)
- [`token.underlying() != u`](https://github.com/sense-finance/sense-v1/blob/dev/pkg/core/src/adapters/abstract/BaseAdapter.sol#L31) (Sense))
- [`ISense(x).maturity() > m`](https://github.com/sense-finance/space-v1/blob/main/src/Space.sol#L79)
- [`ISense(x).swapUnderlyingForPTs(s, m, lent, r);`](https://github.com/sense-finance/sense-v1/blob/3c4335f7fad5609b5c4afeab5a230759930f46da/pkg/core/src/Periphery.sol#L158)
- [`IAPWineToken(principal).getUnderlyingOfIBTAddress()`](https://github.com/APWine/apwine-smart-contracts-public/blob/ec7468cd879bb245cb0ba2881e9df9141b8e80a3/amm/contracts/AMM.sol#L811)
- [`uint256 returned = IAPWineRouter(pool).swapExactAmountIn(i, 1, lent, 0, r, address(this));`](https://github.com/APWine/apwine-smart-contracts-public/blob/ec7468cd879bb245cb0ba2881e9df9141b8e80a3/amm/contracts/AMM.sol#L329)
- [`(IErc20 underlying, ) = token.getUnderlyingToken();`](https://github.com/notional-finance/wrapped-fcash/blob/8f76be58dda648ea58eef863432c14c940e13900/contracts/wfCashBase.sol#L124) (Notional)
- [`token.getMaturity() > m`](https://github.com/notional-finance/wrapped-fcash/blob/8f76be58dda648ea58eef863432c14c940e13900/contracts/wfCashBase.sol#L83) (Notional)
- [`uint256 returned = token.deposit(a - fee, address(this));`](https://github.com/notional-finance/wrapped-fcash/blob/019cfa20369d5e0d9e7a38fea936cc649704780d/contracts/wfCashERC4626.sol#L169) (Notional)
- [`uint128 returned = IYield(y).sellBasePreview(Cast.u128(a));`](https://github.com/yieldprotocol/yieldspace-v2/blob/17b18e26de37e52e504b82f2883cf90249b7a9f5/contracts/Pool.sol#L409)
- [`IYield(y).sellBase(address(this), returned);`](https://github.com/yieldprotocol/yieldspace-v2/blob/17b18e26de37e52e504b82f2883cf90249b7a9f5/contracts/Pool.sol#L369)
Libraries: Element (internal), MarketPlace (internal), Swivel (internal)

### Redeemer

- [`IAPWine(apwineAddr).withdraw(o, amount);`](https://github.com/APWine/apwine-smart-contracts-public/blob/6c98d14464df66ead6c1dee310f9ab9e7a612969/protocol/contracts/protocol/Controller.sol)
- [`ITempus(tempusAddr).redeemToBacking(o, amount, 0, address(this));`](https://github.com/tempus-finance/tempus-protocol/blob/f431e821c81d8cdeae8ad433d160230563f121de/contracts/TempusController.sol)
- [`ISwivel(swivelAddr).redeemZcToken(u, m, amount);`](https://github.com/Swivel-Finance/swivel/blob/main/contracts/v2/swivel/Swivel.sol#L481)
- [`IElementToken(principal).withdrawPrincipal(amount, marketPlace);`](https://github.com/element-fi/elf-contracts/blob/885666433894c598223ea6e32f8cf38236efc2f1/contracts/Tranche.sol)
- [`IYieldToken(principal).redeem(address(this), address(this), amount);`](TODO) // TODO: Get this one.
- [`INotional(principal).maxRedeem(address(this));`](https://github.com/notional-finance/wrapped-fcash/blob/019cfa20369d5e0d9e7a38fea936cc649704780d/contracts/wfCashERC4626.sol#L90)
- [`IPendle(pendleAddr).redeemAfterExpiry(i, u, m);`](https://github.com/pendle-finance/pendle-core/blob/b34d265e4fe8e3a6f79bdec1ab88ab2fd49a882c/contracts/core/PendleRouter.sol)
- [`ISense(d).redeem(o, m, amount);`](https://github.com/sense-finance/sense-v1/blob/3c4335f7fad5609b5c4afeab5a230759930f46da/pkg/core/src/Divider.sol#L305)
Libraries: MarketPlace (internal), Safe (internal)

### MarketPlace

No external contracts used.

## Libraries

Below are the libraries used by Lender.sol, Redeemer.sol and MarketPlace.sol. All libraries are internal and can be located within the repo.

### Lender

Libraries: Element.sol (internal), MarketPlace.sol (internal), Swivel.sol (internal)

### Redeemer

MarketPlace.sol (internal), Safe.sol (internal)
### MarketPlace

Safe.sol (internal)

## **Lender:**

Lender.sol facilitates all lending via the overloaded `lend` method. Users must find the best principal to lend off-chain and then call the right `lend` method and provide the correct parameters to lend capital in the market of their choice. In exchange for lending on Illuminate, the user will receive meta-principal tokens in the market they lent in.

The admin is given privilages to pause principals from lending. This may be important in the event of protocol insolvency or a bug. In addition, the admin is able to able to withdraw fees.

The lender also plays a part in kicking off markets. It is responsible for approving the spending of principal tokens after the creation of the market.

### Areas of Concern

- The admin must not be able to withdraw more fees than what he is entitled to.
- The fee calculation is correct.
- There are no adverse effects from calling the wrong principal for a given `lend` function to other users.
- An appropriate amount of meta-principal tokens are minted to the user whenever a `lend` operation is called.
- The approval process is sound. Once `createMarket` is executed, users are not able to mint meta-principal tokens until `approve` has approved the Lender for all principal tokens in that market.

## **Marketplace:**

Marketplace.sol acts as the central hub for creating new `markets` (defined as an asset-matury pair). The `markets` mapping is stored and admins are the only ones that can create markets.

In addition, MarketPlace.sol will route swaps between the meta-principal token and its respective underlying token via YieldSpace pool. This is tracked via the `pools` mapping that can only be changed by the admin.

Lastly, the MarketPlace mints new meta-principal tokens wheneveer a new market is created. These tokens will conform to the ERC5095 interface.

### Areas of Concern

- The admin cannot withdraw funds without waiting for the period defined by `HOLD`.
- Users are able to swap between their meta-principal tokens and the underlying using YieldSpace pools.
- The admin can add another principal to a market after `createMarket` has been called.

## **Redeemer:**

A redemption is a two step process. First, principal tokens held by the Lender contract must be redeemed for underlying tokens. Second, users must burn their meta-principal tokens to receive their owed underlying assets via Illuminate's `redeem` function.

Redemptions require that a user has the meta-principal token. In addition, to a user manually redeeming their tokens, the meta-principal token contracts can use `authRedeem` to redeem tokens to their users on their behalf.

### Areas of Concern

- Ensure that something out-of-order would not enable a user to redeem more or less tokens than they are entitled to. For example, if a market was partially redeemed (i.e. some principals had not been redeemed by the Redeemer yet), it would not adversely affect a user to call redeem on Illuminate with their meta-principal tokens.
- The `authorized` methods could not be used to incorrectly redeem funds to another user or steal funds in some manner.
