package verify_code

//func TestCreateCode(t *testing.T) {
//	vc := new(VerifyCodeSvc)
//	fmt.Println(generateCode())
//}

//func TestSendMail(t *testing.T) {
//	sender := NewEmailSender(
//		"smtp.qq.com",
//		465,
//		"358110999@qq.com",
//		"lgtgaihorkhnbjdc",
//	)
//
//	go func() {
//		err := main2.Main3()
//		if err != nil {
//			return
//		}
//	}()
//
//	time.Sleep(5 * time.Second)
//	vc := NewVerifyCodeSvc(sender)
//
//	err := vc.SendCodeByEmail(context.Background(),
//		"UserRegister", "用户注册",
//		"358110999@qq.com", "358110999@qq.com")
//	if err != nil {
//		return
//	}
//}
