package usecase

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"math/big"
	"metamask-auth/config"
	"metamask-auth/internal/contract"
)

type contractUseCase struct {
	logger *zap.Logger
	cfg    *config.Config
}

func NewUseCase(logger *zap.Logger, cfg *config.Config) contract.UseCase {
	return &contractUseCase{
		logger: logger,
		cfg:    cfg,
	}
}

func (u *contractUseCase) CreateGame(ctx context.Context) (int64, error) {
	address := common.HexToAddress(u.cfg.Contract.ModeratorAddress)

	privateKey, err := crypto.HexToECDSA(u.cfg.Contract.ModeratorPK)
	if err != nil {
		u.logger.Error("CreateGame err", zap.Error(err))
		return 0, err
	}

	ethereumClient, err := ethclient.Dial(u.cfg.Contract.Rpc)
	if err != nil {
		u.logger.Error("CreateGame err", zap.Error(err))
		return 0, err
	}

	nonce, err := ethereumClient.PendingNonceAt(ctx, address)
	if err != nil {
		u.logger.Error("CreateGame err", zap.Error(err))
		return 0, err
	}

	gasPrice, err := ethereumClient.SuggestGasPrice(ctx)
	if err != nil {
		u.logger.Error("CreateGame err", zap.Error(err))
		return 0, err
	}

	//gasTip, err := ethereumClient.SuggestGasTipCap(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//estimateGas, err := ethereumClient.EstimateGas(ctx, ethereum.CallMsg{
	//	To:       &address,
	//	GasPrice: gasPrice,
	//	Value:    big.NewInt(0),
	//	Data:     []byte{},
	//},
	//)

	moderator, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(u.cfg.Contract.ChainID))
	if err != nil {
		u.logger.Error("CreateGame err", zap.Error(err))
		return 0, err
	}

	myContract, err := contract.NewContract(common.HexToAddress(u.cfg.Contract.ContractAddress), ethereumClient)
	if err != nil {
		u.logger.Error("CreateGame err", zap.Error(err))
		return 0, err
	}

	counter, err := myContract.Counter(&bind.CallOpts{Context: ctx})
	if err != nil {
		u.logger.Error("CreateGame err", zap.Error(err))
		return 0, err
	}

	tx, err := myContract.CreateGame(&bind.TransactOpts{
		From:     address,
		Nonce:    big.NewInt(int64(nonce)),
		Signer:   moderator.Signer,
		Value:    big.NewInt(0),
		GasPrice: gasPrice,
		GasLimit: 100000,
		Context:  ctx,
		NoSend:   true,
	})
	if err != nil {
		u.logger.Error("CreateGame err", zap.Error(err))
		return 0, err
	}

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(u.cfg.Contract.ChainID)), privateKey)
	if err != nil {
		u.logger.Error("CreateGame err", zap.Error(err))
		return 0, err
	}

	err = ethereumClient.SendTransaction(ctx, signedTx)
	if err != nil {
		u.logger.Error("CreateGame err", zap.Error(err))
		return 0, err
	}

	return counter.Int64() + 1, nil
}

func (u *contractUseCase) SetWinner(ctx context.Context, params *contract.WinnerGame) error {
	address := common.HexToAddress(u.cfg.Contract.ModeratorAddress)

	privateKey, err := crypto.HexToECDSA(u.cfg.Contract.ModeratorPK)
	if err != nil {
		return err
	}

	ethereumClient, err := ethclient.Dial(u.cfg.Contract.Rpc)
	if err != nil {
		return err
	}

	nonce, err := ethereumClient.PendingNonceAt(ctx, address)
	if err != nil {
		return err
	}

	gasPrice, err := ethereumClient.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	moderator, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(u.cfg.Contract.ChainID))
	if err != nil {
		return err
	}

	myContract, err := contract.NewContract(common.HexToAddress(u.cfg.Contract.ContractAddress), ethereumClient)
	if err != nil {
		return err
	}

	tx, err := myContract.SetWinner(&bind.TransactOpts{
		From:     address,
		Nonce:    big.NewInt(int64(nonce)),
		Signer:   moderator.Signer,
		Value:    big.NewInt(0),
		GasPrice: gasPrice,
		GasLimit: 100000,
		Context:  ctx,
		NoSend:   true,
	}, big.NewInt(params.ID), common.HexToAddress(params.Receiver))
	if err != nil {
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(u.cfg.Contract.ChainID)), privateKey)
	if err != nil {
		return err
	}

	err = ethereumClient.SendTransaction(ctx, signedTx)
	if err != nil {
		return err
	}

	return nil
}
