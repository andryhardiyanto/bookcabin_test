# BookCabin Frontend

Frontend application for crew voucher generation system built with React.

## Prerequisites

Make sure you have installed:
- Node.js (version 14 or higher)
- npm or yarn
- Backend API server running

## Environment File Setup

1. Create `.env` file in root directory:
```bash
cp .env.example .env
```

2. Edit `.env` file and configure backend host:
```
REACT_APP_API_HOST=http://localhost:8080
```
## Running Instructions

1. Install dependencies:
```bash
npm install
```

2. Start development server:
```bash
npm start
```

3. Open browser at `http://localhost:3000`

Make sure the backend API server is running before using the application.