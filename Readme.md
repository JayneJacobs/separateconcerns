# Separation of concerns
Credits: [Exploring The Go Programming Language(https://www.udemy.com/course/learn-golang/)] Tod McLeod and Daniel Orlando



Separate Code into Three areas of Concern. In this example:

- Access
- Business
- Action

## Description

### Access

<p>
    The access layer allows an user to enter commands. It passes requsts to the Business LogicThe response is relayed to the access layer. This layer also perfoms user authorization. 
</p>

### Business 

<p>
Any functionlity will be defined such as input validation, routing, and logic. The user is created or verified within this layer before passing on a requst. 
This layer is often referred to as middle-ware
</p>


## Action
Actions are defined as the following:
- Database interaction
- logging
- All backend server functions


git tag v0.0.1
separateconcerns 
git push --tags