package tests

//func TestProcessTransactionWhenItsIsValid(t *testing.T) {
//	input := TransactionDtoInput{
//		ID:        "1",
//		AccountID: "1",
//		Amount:    200,
//	}
//
//	expectedOutput := TransactionDtoOutput{
//		ID:           "1",
//		Status:       "approved",
//		ErrorMessage: "",
//	}
//
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//	repositoryMock := mock_entity.NewMockTransactionRepository(ctrl)
//	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Amount, "approved", "").Return(nil)
//
//	usecase := NewProcessTransaction(repositoryMock)
//	output, err := usecase.Execute(input)
//
//	assert.Nil(t, err)
//	assert.Equal(t, expectedOutput, output)
//
//	ctrl.Finish()
//}
