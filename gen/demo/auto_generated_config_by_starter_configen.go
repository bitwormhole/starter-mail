// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package demo

import (
	mail0xcd88fb "github.com/bitwormhole/starter-mail/mail"
	demo10x334aa8 "github.com/bitwormhole/starter-mail/src/demo/golang/demo1"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: com0-demo10x334aa8.Demo1
	cominfobuilder.Next()
	cominfobuilder.ID("com0-demo10x334aa8.Demo1").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDemo1{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDemo1 : the factory of component: com0-demo10x334aa8.Demo1
type comFactory4pComDemo1 struct {

    mPrototype * demo10x334aa8.Demo1

	
	mFromAddrSelector config.InjectionSelector
	mToAddrSelector config.InjectionSelector
	mSenderSelector config.InjectionSelector

}

func (inst * comFactory4pComDemo1) init() application.ComponentFactory {

	
	inst.mFromAddrSelector = config.NewInjectionSelector("${demo.send.from.addr}",nil)
	inst.mToAddrSelector = config.NewInjectionSelector("${demo.send.to.addr}",nil)
	inst.mSenderSelector = config.NewInjectionSelector("#mail.Sender",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDemo1) newObject() * demo10x334aa8.Demo1 {
	return & demo10x334aa8.Demo1 {}
}

func (inst * comFactory4pComDemo1) castObject(instance application.ComponentInstance) * demo10x334aa8.Demo1 {
	return instance.Get().(*demo10x334aa8.Demo1)
}

func (inst * comFactory4pComDemo1) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDemo1) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDemo1) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDemo1) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComDemo1) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDemo1) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.FromAddr = inst.getterForFieldFromAddrSelector(context)
	obj.ToAddr = inst.getterForFieldToAddrSelector(context)
	obj.Sender = inst.getterForFieldSenderSelector(context)
	return context.LastError()
}

//getterForFieldFromAddrSelector
func (inst * comFactory4pComDemo1) getterForFieldFromAddrSelector (context application.InstanceContext) string {
    return inst.mFromAddrSelector.GetString(context)
}

//getterForFieldToAddrSelector
func (inst * comFactory4pComDemo1) getterForFieldToAddrSelector (context application.InstanceContext) string {
    return inst.mToAddrSelector.GetString(context)
}

//getterForFieldSenderSelector
func (inst * comFactory4pComDemo1) getterForFieldSenderSelector (context application.InstanceContext) mail0xcd88fb.Sender {

	o1 := inst.mSenderSelector.GetOne(context)
	o2, ok := o1.(mail0xcd88fb.Sender)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com0-demo10x334aa8.Demo1")
		eb.Set("field", "Sender")
		eb.Set("type1", "?")
		eb.Set("type2", "mail0xcd88fb.Sender")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}




