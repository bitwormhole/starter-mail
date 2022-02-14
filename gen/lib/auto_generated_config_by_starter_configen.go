// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package lib

import (
	support0x5608ca "github.com/bitwormhole/starter-mail/mail/support"
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

	// component: mail.Sender
	cominfobuilder.Next()
	cominfobuilder.ID("mail.Sender").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComSMTPSender{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComSMTPSender : the factory of component: mail.Sender
type comFactory4pComSMTPSender struct {

    mPrototype * support0x5608ca.SMTPSender

	
	mNeedWorkAroundSelector config.InjectionSelector
	mSenderAddressSelector config.InjectionSelector
	mSenderNameSelector config.InjectionSelector
	mSMTPHostSelector config.InjectionSelector
	mSMTPPortSelector config.InjectionSelector
	mSMTPUserSelector config.InjectionSelector
	mSMTPPasswordSelector config.InjectionSelector

}

func (inst * comFactory4pComSMTPSender) init() application.ComponentFactory {

	
	inst.mNeedWorkAroundSelector = config.NewInjectionSelector("${mail.smtp.workaround}",nil)
	inst.mSenderAddressSelector = config.NewInjectionSelector("${mail.sender.address}",nil)
	inst.mSenderNameSelector = config.NewInjectionSelector("${mail.sender.name}",nil)
	inst.mSMTPHostSelector = config.NewInjectionSelector("${mail.smtp.host}",nil)
	inst.mSMTPPortSelector = config.NewInjectionSelector("${mail.smtp.port}",nil)
	inst.mSMTPUserSelector = config.NewInjectionSelector("${mail.smtp.user}",nil)
	inst.mSMTPPasswordSelector = config.NewInjectionSelector("${mail.smtp.password}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComSMTPSender) newObject() * support0x5608ca.SMTPSender {
	return & support0x5608ca.SMTPSender {}
}

func (inst * comFactory4pComSMTPSender) castObject(instance application.ComponentInstance) * support0x5608ca.SMTPSender {
	return instance.Get().(*support0x5608ca.SMTPSender)
}

func (inst * comFactory4pComSMTPSender) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComSMTPSender) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComSMTPSender) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComSMTPSender) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSMTPSender) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSMTPSender) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.NeedWorkAround = inst.getterForFieldNeedWorkAroundSelector(context)
	obj.SenderAddress = inst.getterForFieldSenderAddressSelector(context)
	obj.SenderName = inst.getterForFieldSenderNameSelector(context)
	obj.SMTPHost = inst.getterForFieldSMTPHostSelector(context)
	obj.SMTPPort = inst.getterForFieldSMTPPortSelector(context)
	obj.SMTPUser = inst.getterForFieldSMTPUserSelector(context)
	obj.SMTPPassword = inst.getterForFieldSMTPPasswordSelector(context)
	return context.LastError()
}

//getterForFieldNeedWorkAroundSelector
func (inst * comFactory4pComSMTPSender) getterForFieldNeedWorkAroundSelector (context application.InstanceContext) bool {
    return inst.mNeedWorkAroundSelector.GetBool(context)
}

//getterForFieldSenderAddressSelector
func (inst * comFactory4pComSMTPSender) getterForFieldSenderAddressSelector (context application.InstanceContext) string {
    return inst.mSenderAddressSelector.GetString(context)
}

//getterForFieldSenderNameSelector
func (inst * comFactory4pComSMTPSender) getterForFieldSenderNameSelector (context application.InstanceContext) string {
    return inst.mSenderNameSelector.GetString(context)
}

//getterForFieldSMTPHostSelector
func (inst * comFactory4pComSMTPSender) getterForFieldSMTPHostSelector (context application.InstanceContext) string {
    return inst.mSMTPHostSelector.GetString(context)
}

//getterForFieldSMTPPortSelector
func (inst * comFactory4pComSMTPSender) getterForFieldSMTPPortSelector (context application.InstanceContext) int {
    return inst.mSMTPPortSelector.GetInt(context)
}

//getterForFieldSMTPUserSelector
func (inst * comFactory4pComSMTPSender) getterForFieldSMTPUserSelector (context application.InstanceContext) string {
    return inst.mSMTPUserSelector.GetString(context)
}

//getterForFieldSMTPPasswordSelector
func (inst * comFactory4pComSMTPSender) getterForFieldSMTPPasswordSelector (context application.InstanceContext) string {
    return inst.mSMTPPasswordSelector.GetString(context)
}




