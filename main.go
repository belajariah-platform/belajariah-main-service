package main

import (
	"belajariah-main-service/config"
	"belajariah-main-service/handler"
	"belajariah-main-service/middleware"
	"belajariah-main-service/repository"
	"belajariah-main-service/usecase"
	"belajariah-main-service/utils"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("BELAJARIAH MAIN SERVICE INITIALIZATION")

	//create gin log
	crt, _ := os.Create("gin-port.log")
	gin.DefaultWriter = io.MultiWriter(crt)

	//get config
	configModel := config.GetConfig()

	//get db config
	db := config.ConnectDB(configModel)

	//initiate repository
	userRepository := repository.InitUserRepository(db)
	testRepository := repository.InitTestRepository(db)
	enumRepository := repository.InitEnumRepository(db)
	emailRepository := repository.InitEmailRepository(db)
	classRepository := repository.InitClassRepository(db)
	quranRepository := repository.InitQuranRepository(db)
	storyRepository := repository.InitStoryRepository(db)
	mentorRepository := repository.InitMentorRepository(db)
	ratingRepository := repository.InitRatingRepository(db)
	packageRepository := repository.InitPackageRepository(db)
	paymentRepository := repository.InitPaymentsRepository(db)
	scheduleRepository := repository.InitScheduleRepository(db)
	exerciseRepository := repository.InitExerciseRepository(db)
	learningRepository := repository.InitLearningRepository(db)
	promotionRepository := repository.InitPromotionRepository(db)
	userClassRepository := repository.InitUserClassRepository(db)
	notificationRepository := repository.InitNotificationRepository(db)
	consultationRepository := repository.InitConsultationRepository(db)
	paymentMethodRepository := repository.InitPaymentMethodRepository(db)
	approvalStatusRepository := repository.InitApprovalStatusRepository(db)
	coachingProgramRepository := repository.InitCoachingProgramRepository(db)
	exerciseReadingRepository := repository.InitExerciseReadingRepository(db)
	userClassHistoryRepository := repository.InitUserClassHistoryRepository(db)
	userExerciseReadingRepository := repository.InitUserExerciseReadingRepository(db)

	//initiate usecase
	enumUsecase := usecase.InitEnumUsecase(enumRepository)
	quranUsecase := usecase.InitQuranUsecase(quranRepository)
	storyUsecase := usecase.InitStoryUsecase(storyRepository)
	classUsecase := usecase.InitClassUsecase(classRepository, packageRepository)
	mentorUsecase := usecase.InitMentorUsecase(mentorRepository)
	ratingUsecase := usecase.InitRatingUsecase(ratingRepository)
	packageUsecase := usecase.InitPackageUsecase(packageRepository)
	scheduleUsecase := usecase.InitScheduleUsecase(scheduleRepository)
	exerciseUsecase := usecase.InitExerciseUsecase(exerciseRepository)
	coachingProgramUsecase := usecase.InitCoachingProgramUsecase(coachingProgramRepository)
	promotionUsecase := usecase.InitPromotionUsecase(configModel, promotionRepository, userClassRepository, paymentRepository)
	emailUsecase := usecase.InitEmailUsecase(configModel, userRepository, enumRepository, emailRepository)
	userUsecase := usecase.InitUserUsecase(emailUsecase, userRepository)
	testUsecase := usecase.InitTestUsecase(testRepository, userClassRepository)
	paymentMethodUsecase := usecase.InitPaymentMethodUsecase(paymentMethodRepository)
	exerciseReadingUsecase := usecase.InitExerciseReadingUsecase(exerciseReadingRepository)
	learningUsecase := usecase.InitLearningUsecase(learningRepository, exerciseReadingRepository)
	userExerciseReadingUsecase := usecase.InitUserExerciseReadingUsecase(userExerciseReadingRepository)
	consultationUsecase := usecase.InitConsultationUsecase(userRepository, enumRepository, consultationRepository, approvalStatusRepository)
	paymentUsecase := usecase.InitPaymentUsecase(configModel, emailUsecase, enumRepository, packageRepository, paymentRepository, scheduleRepository, userClassRepository, approvalStatusRepository, userClassHistoryRepository)
	userClassUsecase := usecase.InitUserClassUsecase(configModel, emailUsecase, enumRepository, promotionRepository, userClassRepository, notificationRepository)

	//initiate handler
	mainHandler := handler.InitMainHandler(configModel)
	userHandler := handler.InitUserHandler(userUsecase)
	testHandler := handler.InitTestHandler(testUsecase)
	enumHandler := handler.InitEnumHandler(enumUsecase)
	quranHandler := handler.InitQuranHandler(quranUsecase)
	storyHandler := handler.InitStoryHandler(storyUsecase)
	classHandler := handler.InitClassHandler(classUsecase)
	mentorHandler := handler.InitMentorHandler(mentorUsecase)
	ratingHandler := handler.InitRatingHandler(ratingUsecase)
	paymentHandler := handler.InitPaymentHandler(paymentUsecase)
	packageHandler := handler.InitPackageHandler(packageUsecase)
	scheduleHandler := handler.InitScheduleHandler(scheduleUsecase)
	exerciseHandler := handler.InitExerciseHandler(exerciseUsecase)
	learningHandler := handler.InitLearningHandler(learningUsecase)
	promotionHandler := handler.InitPromotionHandler(promotionUsecase)
	userClassHandler := handler.InitUserClassHandler(userClassUsecase)
	consultationHandler := handler.InitConsultationHandler(consultationUsecase)
	paymentMethodHandler := handler.InitPaymentMethodHandler(paymentMethodUsecase)
	coachingProgramHandler := handler.InitCoachingProgramHandler(coachingProgramUsecase)
	exerciseReadingHandler := handler.InitExerciseReadingHandler(exerciseReadingUsecase)
	userExerciseReadingHandler := handler.InitUserExerciseReadingHandler(userExerciseReadingUsecase)

	//initiate router
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use()
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())
	// router.Use(middleware.AuthMiddleware())

	//router mapping request
	router.GET("/", mainHandler.NoRoute)
	router.GET("/log", mainHandler.Log)

	// router-user
	router.POST("/login", userHandler.LoginUser)
	router.PUT("/user", userHandler.UpdateProfile)
	router.GET("/user/:email", userHandler.GetUser)
	router.POST("/register", userHandler.RegisterUser)
	router.POST("/google_login", userHandler.GoogleLogin)
	router.PUT("/verify_email", userHandler.VerifyEmail)
	router.PUT("/verify_account", userHandler.VerifyAccount)
	router.GET("/check_email/:email", userHandler.CheckEmail)
	router.PUT("/change_password_public", userHandler.ChangePasswordPublic)
	router.PUT("/change_password_private", userHandler.ChangePasswordPrivate)
	router.PUT("/reset_verification", userHandler.ResetVerificationUser)

	// router-test
	router.GET("/tests", testHandler.GetAllTest)
	router.PUT("/test", testHandler.CorrectionTest)

	// router-enum
	router.GET("/enums", enumHandler.GetAllEnum)

	// router-class
	router.GET("/classes", classHandler.GetAllClass)

	// router-story
	router.GET("/stories", storyHandler.GetAllStory)

	// router-mentor
	router.GET("/mentors", mentorHandler.GetAllMentor)
	router.GET("/mentor/:email", mentorHandler.GetMentor)
	router.POST("/mentors", mentorHandler.Mentor)

	// router-rating
	router.POST("/rating_class", ratingHandler.GiveRatingClass)
	router.GET("/rating_class", ratingHandler.GetAllRatingClass)
	router.POST("/rating_mentor", ratingHandler.GiveRatingMentor)

	// router-quran
	router.GET("/qurans", quranHandler.GetAllQuranView)
	router.GET("/quran/surat", quranHandler.GetAllQuran)
	router.GET("/quran/ayat", quranHandler.GetAllAyatQuran)

	// router-package
	router.GET("/packages", packageHandler.GetAllPackage)

	//router-benefit
	router.GET("/benefits", packageHandler.GetAllBenefit)

	// router-payment
	router.GET("/payments", paymentHandler.GetAllPayment)
	router.GET("/payments_reject", paymentHandler.GetAllPaymentRejected)
	router.GET("/payment", paymentHandler.GetAllPaymentByUserID)
	router.POST("/payment", paymentHandler.InsertPayment)
	router.PUT("/payment/confirm", paymentHandler.ConfirmPayment)
	router.PUT("/payment/upload", paymentHandler.UploadPayment)

	// router-exercise
	router.GET("/exercises", exerciseHandler.GetAllExercise)

	// router-exercise-reading
	router.GET("/exercise_reading", exerciseReadingHandler.GetAllExerciseReading)

	// router-exercise-reading
	router.POST("/user_exercise_reading", userExerciseReadingHandler.InserteUserExerciseReading)

	// router-learning
	router.GET("/learnings", learningHandler.GetAllLearning)

	//

	// router-promotion
	router.GET("/promotions", promotionHandler.GetAllPromotion)
	router.GET("/promotion/:code", promotionHandler.GetPromotion)
	router.POST("/promotion/claim", promotionHandler.ClaimPromotion)

	// router-userClass
	router.GET("/user_class", userClassHandler.GetAllUserClass)
	router.GET("/user_class/detail/:code", userClassHandler.GetUserClass)
	router.PUT("/user_class/progress", userClassHandler.UpdateUserClassProgress)

	// router-consultation
	router.GET("/consultations", consultationHandler.GetAllConsultation)
	router.GET("/consultation/user", consultationHandler.GetAllConsultationUser)
	router.GET("/consultation/limit", consultationHandler.GetAllConsultationLimit)
	router.GET("/consultation/mentor", consultationHandler.GetAllConsultationMentor)
	router.GET("/consultation/spam_user", consultationHandler.CheckConsultationSpamUser)
	router.GET("/consultation/spam_mentor", consultationHandler.CheckConsultationSpamMentor)
	router.PUT("/consultation", consultationHandler.UpdateConsultation)
	router.POST("/consultation", consultationHandler.InsertConsultation)
	router.PUT("/consultation/read", consultationHandler.ReadConsultation)
	router.PUT("/consultation/confirm", consultationHandler.ConfirmConsultation)

	// router-payment-method
	router.GET("/payment_methods", paymentMethodHandler.GetAllPaymentMethod)

	// router-payment-method
	router.POST("/coaching_program", coachingProgramHandler.CoachingProgram)

	// router-schedule
	router.POST("/schedule", scheduleHandler.Schedule)

	// go paymentUsecase.CheckAllPaymentExpired()
	// go paymentUsecase.CheckAllPayment2HourBeforeExpired()
	// go promotionUsecase.CheckAllPromotionExpired()
	// go consultationUsecase.CheckAllConsultationExpired()
	// go userClassUsecase.CheckAllUserClassExpired()
	// go userClassUsecase.CheckAllUserClass1DaysBeforeExpired()
	// go userClassUsecase.CheckAllUserClass2DaysBeforeExpired()
	// go userClassUsecase.CheckAllUserClass5DaysBeforeExpired()
	// go userClassUsecase.CheckAllUserClass7DaysBeforeExpired()

	utils.PushLogf("BELAJARIAH MAIN SERVICE STARTED")
	fmt.Println(fmt.Sprintf("BELAJARIAH MAIN SERVICE STARTED ON PORT %d", configModel.Server.Port))

	router.Run(fmt.Sprintf(":%v", configModel.Server.Port))
}
