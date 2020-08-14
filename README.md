# scheduler


  
## How to use
```go
// Set "00:00:00" not "24:00:00".
scheduler.Every().Second(3).From("12:05:03").Do(task)  // Run task every 3 seconds after 12:05:03. That mean it won't run task from 00:00:00 to 12:05:03.
scheduler.Every().Minute(2).From("12:05:03").Do(task)  // Run task every 2 minutes after 12:05:03.
scheduler.Every().Hour(1).From("12:05:03").Do(task)  // Run task every hour after 12:05:03.
// If doesn't set From(), it will set "00:00:00".
  
scheduler.Every().Day(2).At("12:05:03").Do(task)  // Run task at 12:05:03 every two days.
scheduler.Every().Week(0).At("12:05:03").Do(task)  // Set Week(0) for every Sunday, Week(1) for every Monday, ...
// Set Month(1) for 1st of every month, Month(2) for 2st of every month, ...
// If Set Month(31) but it only has 30th this month, the task will be run at 1st of next month.
scheduler.Every().Month(1).At("12:05:03").Do(task)
// If doesn't set At(), the run time will be set as now.
```
  