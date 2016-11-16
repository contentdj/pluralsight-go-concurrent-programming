package main

import (
	"errors"
	"fmt"
)

func main() {
	order := new(PurchaseOrder)
	order.Value = 42.42

	SavePO(order, false).Then(
		func(obj interface{}) error {
			order := obj.(*PurchaseOrder)

			fmt.Printf("Purchase order saved with Id: %d\n", order.Number)

			return errors.New("First promise failed")
		},
		func(err error) {
			fmt.Printf("Failed to save purchae order: %s\n", err.Error())
		}).Then(
		func(obj interface{}) error {
			fmt.Printf("Second promise success\n")

			return nil
		},
		func(err error) {
			fmt.Printf("Failed to save purchae order: %s\n", err.Error())
		})

	fmt.Scanln()
}

type PurchaseOrder struct {
	Number int
	Value  float64
}

func SavePO(order *PurchaseOrder, shoudFail bool) *Promise {
	result := new(Promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		if shoudFail {
			result.failureChannel <- errors.New("Failed to save purchase order")
		} else {
			order.Number = 1234
			result.successChannel <- order
		}
	}()

	return result
}

type Promise struct {
	successChannel chan interface{}
	failureChannel chan error
}

func (promise *Promise) Then(success func(interface{}) error, failure func(error)) *Promise {
	result := new(Promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		select {
		case obj := <-promise.successChannel:
			newErr := success(obj)
			if newErr == nil {
				result.successChannel <- obj
			} else {
				result.failureChannel <- newErr
			}
		case err := <-promise.failureChannel:
			failure(err)
			result.failureChannel <- err
		}
	}()

	return result
}
