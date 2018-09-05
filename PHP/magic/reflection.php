<?php

function getPropertyNames(array $filter = null)
{
    $rc = new ReflectionObject($this);
    $names = array();

    while ($rc instanceof ReflectionClass) {
        foreach ($rc->getPoperties() as $prop) {
            if (!$filter || !in_array($prop->getName(), $filter)) {
                $names[] = $prop->getName();
            }
        }

        $rc = $rc->getParentClass();
    }

    return $names;
}