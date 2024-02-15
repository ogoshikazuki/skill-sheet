// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package entity

import (
	"context"
	"sync"
)

// Ensure, that TransactionControllerMock does implement TransactionController.
// If this is not the case, regenerate this file with moq.
var _ TransactionController = &TransactionControllerMock{}

// TransactionControllerMock is a mock implementation of TransactionController.
//
//	func TestSomethingThatUsesTransactionController(t *testing.T) {
//
//		// make and configure a mocked TransactionController
//		mockedTransactionController := &TransactionControllerMock{
//			TransactionFunc: func(ctx context.Context, f func(tx Tx) error) error {
//				panic("mock out the Transaction method")
//			},
//		}
//
//		// use mockedTransactionController in code that requires TransactionController
//		// and then make assertions.
//
//	}
type TransactionControllerMock struct {
	// TransactionFunc mocks the Transaction method.
	TransactionFunc func(ctx context.Context, f func(tx Tx) error) error

	// calls tracks calls to the methods.
	calls struct {
		// Transaction holds details about calls to the Transaction method.
		Transaction []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F func(tx Tx) error
		}
	}
	lockTransaction sync.RWMutex
}

// Transaction calls TransactionFunc.
func (mock *TransactionControllerMock) Transaction(ctx context.Context, f func(tx Tx) error) error {
	if mock.TransactionFunc == nil {
		panic("TransactionControllerMock.TransactionFunc: method is nil but TransactionController.Transaction was just called")
	}
	callInfo := struct {
		Ctx context.Context
		F   func(tx Tx) error
	}{
		Ctx: ctx,
		F:   f,
	}
	mock.lockTransaction.Lock()
	mock.calls.Transaction = append(mock.calls.Transaction, callInfo)
	mock.lockTransaction.Unlock()
	return mock.TransactionFunc(ctx, f)
}

// TransactionCalls gets all the calls that were made to Transaction.
// Check the length with:
//
//	len(mockedTransactionController.TransactionCalls())
func (mock *TransactionControllerMock) TransactionCalls() []struct {
	Ctx context.Context
	F   func(tx Tx) error
} {
	var calls []struct {
		Ctx context.Context
		F   func(tx Tx) error
	}
	mock.lockTransaction.RLock()
	calls = mock.calls.Transaction
	mock.lockTransaction.RUnlock()
	return calls
}

// Ensure, that TxMock does implement Tx.
// If this is not the case, regenerate this file with moq.
var _ Tx = &TxMock{}

// TxMock is a mock implementation of Tx.
//
//	func TestSomethingThatUsesTx(t *testing.T) {
//
//		// make and configure a mocked Tx
//		mockedTx := &TxMock{
//		}
//
//		// use mockedTx in code that requires Tx
//		// and then make assertions.
//
//	}
type TxMock struct {
	// calls tracks calls to the methods.
	calls struct {
	}
}