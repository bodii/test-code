<?php

include_once "../display_errors.php";

class MobileSniffer
{
  private $userAgent;
  private $device;
  private $browser;
  private $deviceLength;
  private $browserLength;

  public function __construct()
  {
    $this->userAgent = $_SERVER['HTTP_USER_AGENT'];
    $this->userAgent = strtolower($this->userAgent);

    $this->device = array('iphone', 'ipad', 'android', 'silk', 'blackberry', 'touch');#,'ubuntu');
    $this->browser = array('firefox', 'chrome', 'opera', 'msie', 'safari', 'blackberry', 'trident');
    $this->deviceLength = count($this->device);
    $this->browserLength = count($this->browser);
  }

  public function findDevice()
  {
    $device = '没有电脑系统相关信息';
    for($uaSniff=0;$uaSniff < $this->deviceLength; $uaSniff++)
    {
      if(strstr($this->userAgent, $this->device[$uaSniff]))
      {
        $device = $this->device[$uaSniff];
      }
    }
    return $device;
  }

  public function findBrowser()
  {
    $browser = '没有浏览器相关信息';
    for($uaSniff=0; $uaSniff < $this->browserLength; $uaSniff++)
    {
      if(strstr($this->userAgent, $this->browser[$uaSniff]))
      {
        $browser =  $this->browser[$uaSniff];
      }
    }
    return $browser;
  }

}


class Client
{
  private $mobSniff;

  public function __construct()
  {
    $this->mobSniff = new MobileSniffer();
    echo "Device = " .$this->mobSniff->findDevice(). "<br/>";
    echo "Browser = " .$this->mobSniff->findBrowser(). "<br/>";
  }
}

$trigger = new Client();
