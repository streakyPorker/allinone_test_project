package com.example.selfrpc;

import java.io.Serializable;

import lombok.Data;

@Data
public class CallStub extends Stub implements Serializable {
    Object[] args;
    Class[] argTypes;
    Class<?> retType;
}
