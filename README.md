# scheduler
```go
Every(1).Second().From("12:05:03")  // Run task from 12:05:03 every second.
Every(1).Minute().From("12:05:03")  // Run task from 12:05:03 every minute.
Every(1).Hour().From("12:05:03")  // Run task from 12:05:03 every hour.
// If doesn't set From(), it will run task from now.
  
Every(1).Day().At("12:05:03")  // Run task at 12:05:03 every day. Set "00:00:00" not "24:00:00".
Every(1).Week(0).At("12:05:03")  // Set Week(0) for every Sunday, Week(1) for every Monday, ...
// Set Month(1) for 1st of every month, Month(2) for 2st of every month, ...
// If Set Month(31) but there doesn't have 31th of next month, the task will be skipped that month.
Every(1).Month(1).At("12:05:03")
// If doesn't set At(), the run time will be set as now.
```
  