package main

type AuditLogs struct {
    Sha string `json:"sha"`
    Message string `json:"message"`
    Author struct {
        Name string `json:"name"`
        Email string `json:"email"`
    } `json:"author"`
    Timestamp string `json:"timestamp"`
}