# Redeemer

This README document explains each of the contracts in this directory and their relavance to the final product.

## Interfaces.sol

LOC: 75
External Contracts: NA
Libraries: NA

This is a conventional interfaces file. We use it to call methods in external contracts, in particular other procotol's redemption methods.

## MarketPlace.sol

LOC: 14
External Contracts: NA
Libraries: NA

This file can be ignored. It used only to provide the `Principals` enum to the Redeemer contract.

## Redeemer.sol

LOC: 154
External Contracts: 
- [`IAPWine(apwineAddr).withdraw(o, amount);`](https://github.com/APWine/apwine-smart-contracts-public/blob/6c98d14464df66ead6c1dee310f9ab9e7a612969/protocol/contracts/protocol/Controller.sol)
- [`ITempus(tempusAddr).redeemToBacking(o, amount, 0, address(this));`](https://github.com/tempus-finance/tempus-protocol/blob/f431e821c81d8cdeae8ad433d160230563f121de/contracts/TempusController.sol)
- [`ISwivel(swivelAddr).redeemZcToken(u, m, amount);`](https://github.com/Swivel-Finance/swivel/blob/main/contracts/v2/swivel/Swivel.sol#L481)
- [`IElementToken(principal).withdrawPrincipal(amount, marketPlace);`](https://github.com/element-fi/elf-contracts/blob/885666433894c598223ea6e32f8cf38236efc2f1/contracts/Tranche.sol)
- [`IYieldToken(principal).redeem(address(this), address(this), amount);`](TODO) // TODO: Get this one.
- [`INotional(principal).maxRedeem(address(this));`](https://github.com/notional-finance/wrapped-fcash/blob/019cfa20369d5e0d9e7a38fea936cc649704780d/contracts/wfCashERC4626.sol#L90)
- [`IPendle(pendleAddr).redeemAfterExpiry(i, u, m);`](https://github.com/pendle-finance/pendle-core/blob/b34d265e4fe8e3a6f79bdec1ab88ab2fd49a882c/contracts/core/PendleRouter.sol)
- [`ISense(d).redeem(o, m, amount);`](https://github.com/sense-finance/sense-v1/blob/3c4335f7fad5609b5c4afeab5a230759930f46da/pkg/core/src/Divider.sol#L305)
Libraries: MarketPlace (internal), Safe (internal)

### Description

This is a critical contract. It will be used by Illuminate admins and users to route owed debt back to users after their loans have matured. 

Users redeem their debt in a two step process. Once a particular market's loans have matured, users or the protocol may call the `redeem` method for that protocol using its `Principals` enum int value (e.g. Swivel = 1, Yield = 2, ...). Those `redeem` calls will transfer the principal tokens from the `Lender.sol` contract to the redeemer contract's address. From there, the redeemer will call in the principal's `redeem` call, which will transfer owed underlying tokens to the redeemer contract. From there, a user can call `redeem` on Illuminate's `redeem` method, and that will transfer the underlying tokens back to their wallet based on their balance of zcTokens for the given market.

#### Specific Concerns

It is very important that a user is not able to withdraw more capital than they have owed to them. In general, when a principal's `redeem` method is called, we transfer all outstanding principal tokens for that market to the redeemer. From there, it is up to the user to `redeem` their owed debt by calling Illuminate's `redeem`.

We also want to be very careful to avoid reentrancy attacks as these contracts call into many external protocols.

# Safe.sol

LOC: 58
External Contracts: NA
Libraries: NA

This is a utility contract. It's primarily used for optimizing ERC20 functionality.