# Factory Pattern

In class-based programming, the factory method pattern is a creational pattern that uses factory methods to deal with the problem of creating objects without having to specify the exact class of the object that will be created. This is done by creating objects by calling a factory method. 

This example will have two interfaces **INotificationFactory** and **ISender** with their respective implementations to display information about these structures:

- **EmailNotification** with **EmailNotificationSender**
- **SMSNotification** with **SMSNotificationSender**
