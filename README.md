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


### **Order Path:**

A market, defined by an underlying asset (`address`) and a maturity (`uint256`) is created by the admin via the MarketPlace contract. From there, the admin approves the Lender for all principal tokens via the Lender contract.

Next, a user can lend capital for that market by calling `lend`. Depending on the principal he passes to `lend`, he will have to provide different parameters. When `lend` is called, his underlying asset is transferred to the Lender. From there, the Lender converts the underlying asset into principal tokens for the chosen principal. Finally, the `lend` operation ends by minting meta-principal tokens to the user. These tokens will represent his loan in this particular market.

Once the loan has matured, the admin will call `redeem` on all principals servicing a given market. These `redeem` methods will receive the owed principal tokens from the Lender contract and call the redemption functions for each principal. At that point, they will be holding the originally lent capital plus interest.

Next, a user can redeem his owed capital by calling `redeem` with Illuminate's principal. This will result in his meta-principal tokens being burned, and a proportional amount of underlying assets being returned to him.

### ** metaPrincipal token (ERC-5095) functionality:**
TODO: Sourabh
When a user initiates a new fixed-yield position on our orderbook, or manually calls `splitUnderlying`, an underlying token is split into zcTokens and nTokens. (the fixed-yield comes from immediately selling nTokens).

A zcToken (standard erc-20 + erc-2612) can be redeemed 1-1 for underlying upon maturity. After maturity, if a user has not redeemed their zcTokens, they begin accruing interest from the deposit in compound. 

An nToken (non-standard contract balance) is a balance within a users `vault`(vault.notional) within VaultTracker.sol. nTokens (notional balance) represent a deposit in an underlying protocol (compound), and accrue the interest from this deposit until maturity. This interest can be redeemed at any time.


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

TODO: Sourabh

### Areas of Concern
## **Marketplace:**

TODO: Sourabh

### Areas of Concern

## **Redeemer:**

TODO: Sourabh

### Areas of Concern
