<?php
class Caller
{
    public $caller;
    public $module;

    function __call($funcname, $args = array())
    {
        $this->setModuleInformation();

        if (is_object($this->caller) && function_exists('call_user_func_array')) {
            $return = call_user_func_array(array(&$this->caller, $funcname), $args);
        }
        else {
            trigger_error('Call to Function with call_user_func_array failed'."\n", E_USER_ERROR);
        }

        $this->unsetModuleInformation();
        return $return;
    }

    public function __construct($callerClassName = false, $callerModuleName = 'Webboard') {
        if ($callerClassName == false) {
            trigger_error('No Classname', E_USER_ERROR);
        }

        $this->module = $callerModuleName;

        if (class_exists($callerClassName)) {
            $this->caller = new $callerClassName();
        }
        else {
            trigger_error('Class Not Exists:\"\n". $callerClassName' . "\n", E_USER_ERROR);
        }

        if (is_object($this->caller)) {
            $this->setModuleInformation();
            if (method_exists($this->caller, '__init')) {
                $this->caller->__init();
            }

            $this->unsetModuleInformation();
        }
        else {
            trigger_error('Caller is no object!', E_USER_ERROR);
        }
    }

    public function __destruct()
    {
        $this->setModuleInformation();

        if (method_exists($this->caller, '__deinit')) {
            $this->caller->__deinit();
        }

        $this->unsetModuleInformation();
    }

    public function __isset($isset) 
    {
        $this->setModuleInformation();

        if (is_object($this->caller)) {
            $return = isset($this->caller->{$isset});
        }
        else {
            trigger_error('Caller is no object!', E_USER_ERROR);
        }

        $this->unsetModuleInformation();
        return $return;
    }

    public function __unset($unset)
    {
        $this->setModuleInformation();

        if (is_object($this->caller)) {
            if (isset($this->caller->{$unset})) {
                unset($this->caller->{$unset});
            }
        }
        else {
            trigger_error('Caller is no object!', E_USER_ERROR);
        }

        $this->unsetModuleInformation();
    }

    public function __set($set, $val)
    {
        $this->setModuleInformation();
        
        if (is_object($this->caller)) {
            $this->caller->{$set} = $val;
        }
        else {
            trigger_error('Caller is no object!', E_USER_ERROR);
        }

        $this->unsetModuleInformation();
    }

    public function __get($get) 
    {
        $this->setModuleInformation();

        if (is_object($this->caller)) {
            if (isset($this->caller->{$get})) {
                $return = $this->caller->{$get};
            }
            else {
                $return = false;
            }
        }
        else {
            trigger_error('Caller is no object!', E_USER_ERROR);
        }

        $this->unsetModuleInformation();
        return $return;
    }

    public function setModuleInformation()
    {
        $this->caller->module = $this->module;
    }

    public function unsetModuleInformation()
    {
        $this->caller->module = null;
    }
}


class Config 
{
    public $module;
    public $test;

    public function __construct()
    {
        print('Constructor will have no Module Information...Use __init() instead!'. "\n");
        print('-->' . print_r($this->module, 1) . '<--');
        print("\n");
        print("\n");
        $this->test = '123'; 

    }

    public function __init()
    {
        print('Using of __init()!'. "\n");
        print('-->' . print_r($this->module, 1) . '<--');
        print("\n");
        print("\n");
    }

    public function testFunction($test = false) 
    {
        if ($test != false) {
            $this->test = $test;
        }
    }
}

echo('<pre>');
$wow = new Caller('Config', 'Guestbook');
print_r($wow->test);
print("\n");
print("\n");
$wow->test = '456';
print_r($wow->test);
print("\n");
print("\n");
$wow->testFunction('789');
print_r($wow->test);
print("\n");
print("\n");
print_r($wow->module);
echo('</pre>');
