# trans
go类型转化工具包

平时工作开发中，会有很多类型转换的工作，本工具实现任意常用类型之间的转换

不用再关心各种类型之间的转换方法，直接调用方法传入任意类型的参数获取自己想要的类型即可。
与cast库相比：支持指针类型，同时解决cast库中的一些转化问题，比如cast库，int类和int64类型的1转到bool类型分别为true和false，
但本库转化之后都是true


