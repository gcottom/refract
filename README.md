# Refract
You've heard of reflection. 

## Refract is built on top of reflect. 
The purpose of refract is to provide a simple layer between reflect and your code.
Reflect is unforgiving. If you make an error in reflect, it likely results in panic. With refract, 
most functions return an error that can be handled gracefully. Refract allows you to use dynamic 
structs in new ways. Use the built in Len, Append, Preppend, and other utility functions to improve workflow. 