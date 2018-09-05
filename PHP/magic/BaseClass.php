<?php

class BaseClass
{
    public $imported;
    public $imported_functions;

    public function __construct()
    {
        $imported = array();
        $imported_functions = array();
    }

    public function imports($object)
    {
        $new_imports = new $object();
        $imports_name = get_class($new_imports);
        array_push($imported, array($imports_name, $new_imports));
        $imports_function = get_class_methods($new_imports);
        foreach ($imports_function as $i=>$function_name) {
            $this->imported_functions[$function_name] = &$new_imports;
        }
    }

    public function __call($method, $args)
    {
        if (array_key_exists($method, $this->imported_functions)) {
            return call_user_func_array(
                array($this->imported_functions[$method], $method),
                $args
            );
        }

        throw new ErrorException(
            'Call to Undefined Method/Class Function',
            0,
            E_ERROR
        );
    }
    
}


class B extends BaseClass
{
    public function __construct()
    {
        $this->imports('ExternalFunc');
    }

    public function Msg()
    {
        echo 'Hello world' . PHP_EOL;
    }
}

// $b = new B();
// $b->Msg();
// $b->TestB();