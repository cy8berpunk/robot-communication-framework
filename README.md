# Robot Communication Framework

The RCF is a framework for data distribution, which is the most essential part of an autonomous platform. It is very similar to [ROS](https://www.ros.org/) but without packages and the C/C++ complexity overhead while still maintaining speed and **safe** thread/ lang. standards, thanks to the [go](https://golang.org/) lang.

# Go-Routine Memory Synchronisation

Since sharing memory is a very complicated and difficult thing to get right without overcomplicating things, Go channels are used, according to Golang's motto "Share memory by communicating, don't communicate by sharing memory"!

# Installation

Installation via. command line: <br>

`go get https://github.com/cy8berpunk/robot-communication-framework` <br>

installation from code

`
import (
    "fmt"
    "github.com/cy8berpunk/robot-communication-framework"
)
`

# Concept

The primary communication interface is a node, in contrast to ROS, or various other robot platforms, the node is only a object instance and does not contain any code. A node resembles the platform for services and topics.

Every node has a port number which also resembles the node ID, but no actual name, so every node is represented through a number. This is a major part of the concept to reduce complexity since there is no need for internal node communication to resolve addresses.

## Topics

Every topic represents a communication channel, from which data can be pulled from or pushed onto or to which can be listened.
The topic communication is split up into elements or msg's, every element/msg represents a byte array pushed to the topic. There are no
variables assignments a element/ msg can represent a single value or anything else. If a topic element/msg structure is needed, the `glob` methods serialize a string map and use the serialized maps as elements/ msgs and as such enable a structured and more generic way to use topics. 
A topic can be identified via its name and the node(node ID) which it is hosted by.

## Services

A service is a function that can be executed by nodes or node clients. Since they are not meant to do calculations but to provide node side
functionality for the clients, they can only be called without a return value.

## Ways to communicate data

A topic is meant to share command & control or sensor data, hence data that needs to be accurate and which does not require high bandwith, since a topic rely's on tcp sockets to communicate.
Data that is more error tolerant, such as video, image or audio data and requires high bandwidth can be shared via streamed topics, which are very similar to normal topics, with the only difference that the data is shared via UDP instead of TCP.    
