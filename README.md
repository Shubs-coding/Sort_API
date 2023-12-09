# Go Sorting Server

This Go application demonstrates sequential and concurrent processing of sorting arrays using Go's concurrency features.

## Features

- Two endpoints for sorting arrays: `/process-single` for sequential processing and `/process-concurrent` for concurrent processing.
- Accepts input in JSON format with an array of sub-arrays to be sorted.
- Measures the time taken to sort arrays in nanoseconds using Go's time package.
- Containerized with Docker for easy deployment.

  
