package service

import "github.com/google/wire"

var (
	ProviderSet = wire.NewSet(
		// 以下两种方式均能注入 RightServiceImpl，然后通过 wire.Bind 绑定到 RightsService 接口

		wire.Struct(new(RightsServiceImpl), "*"),
		// NewRightsService,

		// 以下两种方式均能将 RightsServiceImpl 绑定到 RightsService 接口；
		// 区别是第一种方法实现了类型之间的绑定，RightsServiceImpl 的依赖也会被 wire 注入；
		// 第二种方法会创建一个 RightsServiceImpl 变量，然后将其注入到需要 RightService 的地方，但是该变量的依赖不会被 wire 注入

		wire.Bind(new(RightsService), new(*RightsServiceImpl)), // 注意，这里不能是 RightsServiceImpl，因为是指针类型实现了接口
		// wire.InterfaceValue(new(RightsService), &RightsServiceImpl{}), // 注意，这里不能是 *RightsServiceImpl

		// 注入普通的 UserService
		// NewUserService,

		// 注入特殊的 UserService：来自于 SpecialRightServiceImpl 的 UserService
		NewSpecialRightService,
		wire.FieldsOf(new(*SpecialRightServiceImpl), "SpecialUserService"),
	)
)
