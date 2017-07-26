<?php 

/* 
    [ 外观设计模式 (Facade Pattern) ]
    
    外观模式又叫门面模式，是为子系统中的一组接口提供一个一致的界面，定义一个高层
    接口，这个接口使得这一子系统更加容易使用。

    ×× 为什么需要外观模式
    1. 开发阶段，子系统越来越复杂，增加外观模式提供一个简单的调用接口。
    2. 维护一个大型遗留系统的时候，可能这个系统已经非常难以维护和扩展，
    但又包含非常重要的功能，为其开发一个外观类，以便新系统与其交互。
    3. 外观模式可以隐藏来自调用对象的的复杂性。

*/

// 实例
/* 比如说我们去医院就诊，医院有医生员工系统，有药品系统，有患者资料系统。
但是我们只是在前台挂个号，就能在其他系统里都看到我们。外观系统就差不多这
样。
*/


// 医院医生员工系统
class DoctorSystem
{
    // 通知就诊医生
    public static function getDoctor($name)
    {
        echo __CLASS__ . ': [' . $name . '] 医生，有患者挂你号', '<br>';
        Doctor::setDoctorName($name);
        return (new Doctor);
    }
}

// 医生类
class Doctor
{
    private static $doctorName;
    private static $sufferName;
    
   public static function setDoctorName($doctorName)
   {
       self::$doctorName = $doctorName;
   }

   public static function setSufferName($sufferName)
   {
       self::$sufferName = $sufferName;
   }

    public function prescribe($data)
    {
        echo __CLASS__ . ': ['. self::$doctorName . '] 医生 开个处方给 [' . self::$sufferName . ']', '<br>';
        return [ 'sufferName'=>self::$sufferName, 'data'=>'祖传秘方， 药到病除'];
    }
}

// 患者系统
class SufferSystem
{
    public static function getData($suffer)
    {
        $data = $suffer . '资料';

        echo __CLASS__ . ':病人 [' . $suffer . '] 的资料是这些'. '<br>';
        Doctor::setSufferName($suffer);
        return $data;
    }
}

// 药房
class shop
{
    public static $medicine;

    public static function setMedicine($medicine)
    {
        self::$medicine = $medicine;
    }

    public static function getMedicine()
    {
        echo __CLASS__ . ':' . self::$medicine . '<br>';
    }
}

// 医药系统
class MedicineSystem
{
    public static function register($prescribe)
    {
        echo __CLASS__ . ': [' . $prescribe['sufferName'] .']患者 拿到处方： '.$prescribe['data'] . '------
        通知药房发药了'. '<br>';
        Shop::setMedicine('药5千克');
    }
}


// 持号系统

class Facade
{
    static public function regist($suffer, $doctor)
    {
        // 通知就诊医生
        $doct = DoctorSystem::getDoctor($doctor);

        // 患者系统拿病历资料
        $data = SufferSystem::getData($suffer);

        // 医生看病历资料，开处方
        $prscirbe = $doct->prescribe($data);

        // 医生系统登记处方
        MedicineSystem::register($prscirbe);

        // 药房拿药
        Shop::getMedicine();
    }
}


Facade::regist('叶好成', '贾一');

/*
    外观模式，也叫门面模式。它多用于多个子系统之间，作为中间层。用户通过Facade
    对象，直接请求工作， 省去了用户调用多个子系统的复杂动作。

    外观模式常举的一个例子，就是我们买了好多支股票，但是时间有限。盯盘很复杂，我
    们搞得一团糟。所以， 我们干脆买了股票基金。股票基金就好比于外观模式的Facade
    对象，而子系统就是股票基金投的各支股票。
*/