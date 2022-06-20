# Lender

This README document explains each of the contracts in this directory and their relavance to the final product.

## Lender.sol

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

External Contracts: NA
Libraries: NA

This is a conventional interfaces file. We use it to call methods in external contracts, in particular other procotol's redemption methods.

## Safe.sol

External Cntracts: NA
Libraries: NA

This is a utility contract. It's primarily used for optimizing ERC20 functionality.

## Element.sol

External Contracts: NA
Libraries: NA

## Swivel.sol

External Contracts: NA
Libraries: NA

## MarketPlace.sol

External Contracts: NA
Libraries: NA

## Cast.sol

External Contracts: NA
Libraries: NA

This is a utility contract. It is used to safely cast between differently sized integer types.

### LOC Report

--------------------------------------------------------------------------------
File                                         blank        comment           code
--------------------------------------------------------------------------------
./test/lender/Lender.sol                       111            234            392
./test/lender/Interfaces.sol                    54            104            137
./test/lender/Safe.sol                          32             47             95
./test/lender/Element.sol                        6              1             22
./test/lender/Swivel.sol                         3              2             19
./test/lender/MarketPlace.sol                    2              2             14
./test/lender/Cast.sol                           2              3              7
--------------------------------------------------------------------------------
SUM:                                           210            393            686
--------------------------------------------------------------------------------

### Description

This is a critical contract. It will be used by Illuminate admins and users to lend capital.

This contract works by heavily overloading the `lend` method. Each lend is also passed a `uint8 p` parameter which represents the protocol to use (based on the `Principals` enum). From there, the user's principal tokens are transferred to the lender contract, which in turn executes the appropriate operation to execute a lend. The user then receives an equivalent amount of zcTokens to later redeem, upon maturity, their original captial plus interest.

At a high level, the lend operations work as follows:
0. Lender verifies that the maturity and underlying are correct for the given principal token
1. User transfers their principal tokens to the lender contract
2. A fee is calculated and the given principal's fee total to withdraw is updated in `fees`
3. The lend operation is executed using the remaining principal tokens
4. Based on the lend operation, a certain amount of zcTokens are minted to the user

After maturity, the user should be able to use the zcTokens to retrieve the owed underlying via the redeemer contract.
#### Specific Concerns

The user must not be able to mint more zcTokens than appropriate for a given loan. 

In addition, the withdrawl method should not be able to extract more fees than the lender contract is entitled to collect.

We also want to be very careful to avoid reentrancy attacks as these contracts call into many external contracts.