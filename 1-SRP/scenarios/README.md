# Single Responsibility Principle (SRP) - Practice Scenarios

This directory contains different scenarios to practice implementing the Single Responsibility Principle in Go.

## Overview

Each scenario is designed to help you understand and apply SRP by breaking down complex systems into focused, single-responsibility components.

## Available Scenarios

### 1. [📧 Email Service](./email_services/README.md)
Build an email service with separate components for:
- Email composition with templates
- Email validation
- Sending emails
- Email logging
- Attachment management

### 2. [📊 Report Generator](./report_generator/README.md)
Create a report generation system that separates:
- Data fetching
- Data processing
- Report formatting
- Report storage
- User notifications

### 3. [🎮 Game Character System](./game_character_system/README.md)
Implement a character system with distinct components for:
- Character attributes management
- Inventory system
- Combat actions
- Movement control
- State persistence

### 4. [🏦 Banking Transaction Processor](./banking_transaction_processor/README.md)
Develop a transaction system with separate responsibilities for:
- Transaction validation
- Transaction processing
- History recording
- Notification handling
- Balance management

### 5. [🖼️ Image Processing Service](./image_processing_service/README.md)
Create an image service with distinct components for:
- Image uploading
- Format validation
- Image processing
- Storage management
- Metadata handling
- Cache management

## How to Use These Scenarios

1. Start with one scenario
2. Identify the different responsibilities
3. Design interfaces for each responsibility
4. Implement concrete types for each interface
5. Create tests for your implementation
6. Review and refactor as needed

## Best Practices

When implementing these scenarios, remember to:
- Create clear interfaces
- Use dependency injection
- Keep components loosely coupled
- Implement proper error handling
- Write unit tests for each component
- Document your code
- Follow Go idioms and conventions

## Getting Started

1. Choose a scenario
2. Create a new directory for your implementation
3. Start with the interfaces
4. Implement the concrete types
5. Write tests
6. Compare with example solutions

## Running the Examples

```bash
cd <scenario-directory>
go test ./...
```