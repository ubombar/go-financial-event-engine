# Go Financial Event Engine
This is a pet project of mine to optimize my financial stuff. The idea is like this;
1. You need to create an `events` file by initializing using the CLI tool.
2. Then you can start adding events to the database. These events can be like; 
   - `Income/Expense Event`. It has a fixed event time and amount referring to the account (Well there can be different accounts with different amounts and assets).
   - `Variable Deposit/Withdraw Event`: The idea is to add a variable to the event amount. There might be different ways to specify possible outcomes. For now, I will only implement the binomial amount with the given probability.
   - `Variable Date Event`: This is a future thought but the events themselves can occur on different dates as well. Right now I will not use this but adding it here in case I ever need it.
3. The engine will create a tree structure in the background for these events. Normal Income/Expense events will only be appended to the tree however the variable ones will fork the tree. 
4. Then at last when you have the tree, you can execute certain conditions and if they hold.

## Future Plans
This looks like an interesting project, I was doing my finances when I came up with this. Let me know if something similar exists by contacting me!

Additionally, I might consider to make this a language as well. I am thinking of a declarative structure to declare the event tree and a query language for the validation.

## Contributions
All ideas and contributions are well welcomed! However, I might not have time to maintain this repo, we will see with time...

