package com.example.allinone2.provider;

import com.example.allinone2.DemoService;
import org.apache.dubbo.config.annotation.DubboService;
import org.springframework.stereotype.Component;

/**
 * @author lzy
 */
@DubboService
public class DemoServiceImpl implements DemoService {
    @Override
    public String hello(String text) {
        return "hello " + text;
    }
}
