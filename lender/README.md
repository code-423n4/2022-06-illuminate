# Lender

This README document explains each of the contracts in this directory and their relavance to the final product.

## Lender.sol

LOC: 358
External Contracts: 
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



## Interfaces.sol

LOC: 121
External Contracts: NA
Libraries: NA

This is a conventional interfaces file. We use it to call methods in external contracts, in particular other procotol's redemption methods.

## Safe.sol

LOC: 95
External Cntracts: NA
Libraries: NA

This is a utility contract. It's primarily used for optimizing ERC20 functionality.

## Element.sol

LOC: 22
External Contracts: NA
Libraries: NA

## Swivel.sol

LOC: 19
External Contracts: NA
Libraries: NA

## MarketPlace.sol

LOC: 14
External Contracts: NA
Libraries: NA

## Cast.sol

LOC: 7
External Contracts: NA
Libraries: NA

This is a utility contract. It is used to safely cast between differently sized integer types.

### Description

This is a critical contract. It will be used by Illuminate admins and users to route owed debt back to users after their loans have matured. 

Users redeem their debt in a two step process. Once a particular market's loans have matured, users or the protocol may call the `redeem` method for that protocol using its `Principals` enum int value (e.g. Swivel = 1, Yield = 2, ...). Those `redeem` calls will transfer the principal tokens from the `Lender.sol` contract to the redeemer contract's address. From there, the redeemer will call in the principal's `redeem` call, which will transfer owed underlying tokens to the redeemer contract. From there, a user can call `redeem` on Illuminate's `redeem` method, and that will transfer the underlying tokens back to their wallet based on their balance of zcTokens for the given market.

#### Specific Concerns

It is very important that a user is not able to withdraw more capital than they have owed to them. In general, when a principal's `redeem` method is called, we transfer all outstanding principal tokens for that market to the redeemer. From there, it is up to the user to `redeem` their owed debt by calling Illuminate's `redeem`.

We also want to be very careful to avoid reentrancy attacks as these contracts call into many external protocols.