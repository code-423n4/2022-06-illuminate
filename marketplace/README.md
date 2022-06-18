# MarketPlace

This README document explains each of the contracts in this directory and their relavance to the final product.

## ERC2612.sol

LOC: 31
External Contracts: NA
Libraries: NA

Implementation of ERC2612 used by the ZcToken.

## Hash.sol

LOC: 45
External Contracts: NA
Libraries: NA

This is a utility contract. It's primarily used for optimizing ERC20 functionality.

## IErc2612.sol

LOC: 31
External Contracts: NA
Libraries: NA

This is in interface that is used by ZcToken.sol.

## Interfaces.sol

LOC: 47
External Contracts: NA
Libraries: NA

This is a standard interfaces file used 

## IPErc20.sol

LOC: 14
External Contracts: NA
Libraries: NA

This is an interface implemented by ERC-2612, which is implemented by the ZcToken.

## IZcToken.sol

LOC: 6
External Contracts: NA
Libraries: NA

This is the interface for the ZcToken, which implements ERC-2612. The purpose of ERC-2612 is to allow contracts to approve spending of tokens. In the case of Illuminate, this is done when the lender approves redeemer usage of ZcTokens.

## PErc20.sol

LOC: 91
External Contracts: NA
Libraries: NA

This file is an implementation of the PERC20 interface, which is implemented by Erc2612.sol.

## Safe.sol

LOC: 47
External Contracts: NA
Libraries: NA

This is a utility contract. It's primarily used for optimizing ERC20 functionality.

## ZcToken.sol

LOC: 34
External Contracts: NA
Libraries: NA

This is the token contract that represents user loan's on Illuminate. When a user makes a loan on Illuminate, they receive ZcTokens that represent the amount of underlying tokens they are owed. 

This contract implements the ERC2612 interface, which implements the PERC20 (permissioned-ERC20) interface. The purpose of this is to allow contracts to approve spending of the tokens. In the case of Illuminate, we require that the lender and redeemer be given access to these tokens to execute lend and redeem operations.

## MarketPlace.sol

LOC: 122
// TODO: Update the interface and add links to the IPool interface
External Contracts: 
- [`return pool.sellPrincipalToken(msg.sender, pool.sellPrincipalTokenPreview(a));`]()
- [`Safe.transfer(IErc20(address(pool.principalToken())), address(pool), a);`]()
- [`return pool.sellUnderlying(msg.sender, pool.sellUnderlyingPreview(a));`]()
- [`Safe.transfer(IErc20(address(pool.principalToken())), address(pool), a);`]()
- [`return pool.buyUnderlying(msg.sender, pool.buyUnderlyingPreview(a), a);`]()

Libraries: Safe (internal)

This file can be ignored. It used only to provide the `Principals` enum to the Redeemer contract.

### Description

This is a critical contract. It will be used by Illuminate to create new lending markets and facilitate swaps of principal tokens for their underlying via YieldSpace pools.

There are two critical functionalities provided by the marketplace contract.

First, there is the creation of markets. A market is defined as the tuple of maturity (defined by a `uint256` timestamp) and an underlying asset (identified by its contract `address`). When creating a market, the `admin` will provide a list of addresses that correspond to each principal's index value in the `Principals` enum. When a `lend` or `redeem` method is called later, the contracts will be able to retrieve those addresses via the `markets` mapping held in the MarketPlace contract.

Second, the MarketPlace contract is used to facilitate swaps between the zcTokens and underlying assets via their respective YieldSpace pools. These are defined in the `pools` mapping and can be added by the `admin` after the creation of the pool. In addition, `pools` may be paused on a principal basis by the `admin`.

#### Specific Concerns

The MarketPlace contract is primarily operated by the `admin`. As such, it is very important that only the `admin` be able to edit the `pools` and `markets` attributes. If either the `pools` or the `markets` attributes were to change, it would present substantial security risks to the Illuminate protocol.
