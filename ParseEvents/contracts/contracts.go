package contracts

import (
	"ParseEvents/config"
	"ParseEvents/contracts/abigen/collection"
	"ParseEvents/contracts/abigen/factory"
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"time"
)

type Parse interface {
	Start(ctx context.Context) error
}

type Events struct {
	cfg *config.Config
}

func NewEvents(
	config *config.Config,
) *Events {
	return &Events{
		cfg: config,
	}
}

func (e *Events) Start(ctx context.Context) error {
	ethereumClient, err := ethclient.Dial(e.cfg.Contract.Rpc)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	eg := new(errgroup.Group)
	eg.Go(func() error {
		return e.watchEventCollectionCreated(ctx, ethereumClient)
	})
	eg.Go(func() error {
		return e.watchEventTokenMinted(ctx, ethereumClient)
	})

	if err = eg.Wait(); err != nil {
		return err
	}

	return nil
}

func (e *Events) watchEventCollectionCreated(ctx context.Context, client *ethclient.Client) error {
	logFile, err := os.OpenFile("event_fac_logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)

		return err
	}
	defer logFile.Close()

	collectionFactoryContract, err := factory.NewFactoryFilterer(common.HexToAddress(e.cfg.Contract.CollectionFactoryAddress), client)
	if err != nil {
		log.Fatalf("NewFactoryFilterer err: %v", err)

		return err
	}

	events := make(chan *factory.FactoryCollectionCreated)
	opts := &bind.WatchOpts{
		Start:   nil,
		Context: ctx,
	}

	subscription, err := collectionFactoryContract.WatchCollectionCreated(opts, events)
	if err != nil {
		log.Fatalf("Failed to subscribe to the CollectionCreated events: %v", err)

		return err
	}
	defer subscription.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return nil
		case errChan := <-subscription.Err():
			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
			return errChan
		case event := <-events:
			j, _ := json.MarshalIndent(
				CollectionCreatedEvent{
					Event:      "CollectionCreated",
					Collection: event.Collection.Hex(),
					Name:       event.Name,
					Symbol:     event.Symbol,
					Timestamp:  time.Now(),
				},
				"",
				"  ",
			)

			jsonData, err := json.MarshalIndent(j, "", "  ")
			if err != nil {
				log.Fatalf("Failed to marshal event to JSON: %v", err)

				return err
			}

			_, err = logFile.WriteString(string(jsonData) + "\n")
			if err != nil {
				log.Fatalf("Failed to write event to file: %v", err)

				return err
			}
		}
	}
}

func (e *Events) watchEventTokenMinted(ctx context.Context, client *ethclient.Client) error {
	logFile, err := os.OpenFile("event_col_logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)

		return err
	}
	defer logFile.Close()

	collectionContract, err := collection.NewCollectionFilterer(common.HexToAddress(e.cfg.Contract.CollectionAddress), client)
	if err != nil {
		log.Fatalf("NewCollectionFilterer err: %v", err)

		return err
	}

	events := make(chan *collection.CollectionTokenMinted)
	opts := &bind.WatchOpts{
		Start:   nil,
		Context: ctx,
	}

	subscription, err := collectionContract.WatchTokenMinted(opts, events)
	if err != nil {
		log.Fatalf("Failed to subscribe to the TokenMinted events: %v", err)

		return err
	}
	defer subscription.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return nil
		case errChan := <-subscription.Err():
			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
			return errChan
		case event := <-events:
			j, _ := json.MarshalIndent(
				TokenMintedEvent{
					Event:      "",
					Collection: event.Collection.Hex(),
					Recipient:  event.Recipient.Hex(),
					TokenId:    event.TokenId,
					TokenURI:   event.TokenURI,
					Timestamp:  time.Now(),
				},
				"",
				"  ",
			)

			jsonData, err := json.MarshalIndent(j, "", "  ")
			if err != nil {
				log.Fatalf("Failed to marshal event to JSON: %v", err)

				return err
			}

			_, err = logFile.WriteString(string(jsonData) + "\n")
			if err != nil {
				log.Fatalf("Failed to write event to file: %v", err)

				return err
			}
		}
	}
}
