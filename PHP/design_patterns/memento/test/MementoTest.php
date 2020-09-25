<?php

namespace design_patterns\memento\test;

use design_patterns\memento\State;
use design_patterns\memento\Ticket;

// 目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入vendor下的autoload文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用vendor下的PHPUnit测试类
use PHPUnit\Framework\TestCase;

/**
 * MemetoTest class
 * 备忘录测试类
 */
class MemetoTest extends TestCase
{
    /**
     * TestOpenTicketAssignAndSetBackToOpen function
     */
    public function testOpenTicketAssignAndSetBackToOpen()
    {
        $ticket = new Ticket();

        $ticket->open();
        $openedState = $ticket->getState();
        $this->assertEquals(State::STATE_OPENED, (string) $ticket->getState());

        $memento = $ticket->saveToMemento();

        $ticket->assign();
        $this->assertEquals(State::STATE_ASSIGNED, (string) $ticket->getState());
    
        $ticket->restoreFromMemento($memento);

        $this->assertEquals(State::STATE_OPENED, (string) $ticket->getState());
        var_dump($openedState);
        var_dump($ticket->getState());
        // 测试不相同：false，这里是相同的
        $this->assertNotSame($openedState, $ticket->getState());

    }
}

