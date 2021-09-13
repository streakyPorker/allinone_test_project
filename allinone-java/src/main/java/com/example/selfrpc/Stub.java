package com.example.selfrpc;

import java.io.Serializable;

import lombok.Data;

/**
 * @author lzy
 */
@Data
public class Stub implements Serializable {
    String serviceName;
    String methodName;

}


