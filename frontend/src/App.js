import React, { useState } from 'react';
import axios from 'axios';
import './App.css';

// Get API host from environment variable
const API_HOST = process.env.REACT_APP_API_HOST || 'http://localhost:8080';

const App = () => {
  const [formData, setFormData] = useState({
    name: '',
    id: '',
    flightNumber: '',
    date: '',
    aircraft: ''
  });
  
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState(false);
  const [generatedSeats, setGeneratedSeats] = useState([]);

  const aircraftTypes = ['ATR', 'Airbus 320', 'Boeing 737 Max'];

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
    // Clear error when user starts typing
    if (error) setError('');
    if (success) setSuccess(false);
  };

  const convertToAPIDate = (dateString) => {
    // Convert from YYYY-MM-DD to YYYY-MM-DD format for API
    return dateString;
  };

  const validateForm = () => {
    if (!formData.name.trim()) {
      setError('Crew name is required');
      return false;
    }
    if (!formData.id.trim()) {
      setError('Crew ID is required');
      return false;
    }
    if (!formData.flightNumber.trim()) {
      setError('Flight number is required');
      return false;
    }
    if (!formData.date) {
      setError('Flight date is required');
      return false;
    }
    if (!formData.aircraft) {
      setError('Aircraft type is required');
      return false;
    }
    return true;
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    if (!validateForm()) return;

    setLoading(true);
    setError('');
    setSuccess(false);

    try {
      const checkResponse = await axios.post(`${API_HOST}/api/check`, {
        flightNumber: formData.flightNumber,
        date: convertToAPIDate(formData.date)
      });

      if (checkResponse.data.data.exists) {
        setError('Vouchers have already been generated for this flight and date.');
        setLoading(false);
        return;
      }

      const generateResponse = await axios.post(`${API_HOST}/api/generate`, {
        name: formData.name,
        id: formData.id,
        flightNumber: formData.flightNumber,
        date: convertToAPIDate(formData.date),
        aircraft: formData.aircraft
      });

      if (generateResponse.data.success) {
        setGeneratedSeats(generateResponse.data.data.seats || []);
        setSuccess(true);
        setFormData({
          name: '',
          id: '',
          flightNumber: '',
          date: '',
          aircraft: ''
        });
      }

    } catch (err) {
      console.error('API Error:', err);
      if (err.response?.data?.message) {
        setError(err.response.data.message);
      } else if (err.code === 'ECONNREFUSED') {
        setError(`Unable to connect to the server. Please ensure the API server is running on ${API_HOST}.`);
      } else {
        setError('An error occurred while processing your request. Please try again.');
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="app">
      <div className="container">
        <div className="header">
          <h1>BookCabin</h1>
          <p>Crew Voucher Generation System</p>
        </div>

        <div className="form-container">
          <form onSubmit={handleSubmit} className="voucher-form">
            <div className="form-group">
              <label htmlFor="name">Crew Name</label>
              <input
                type="text"
                id="name"
                name="name"
                value={formData.name}
                onChange={handleInputChange}
                placeholder="Enter crew name"
                disabled={loading}
              />
            </div>

            <div className="form-group">
              <label htmlFor="id">Crew ID</label>
              <input
                type="text"
                id="id"
                name="id"
                value={formData.id}
                onChange={handleInputChange}
                placeholder="Enter crew ID"
                disabled={loading}
              />
            </div>

            <div className="form-group">
              <label htmlFor="flightNumber">Flight Number</label>
              <input
                type="text"
                id="flightNumber"
                name="flightNumber"
                value={formData.flightNumber}
                onChange={handleInputChange}
                placeholder="Enter flight number (e.g., GA102)"
                disabled={loading}
              />
            </div>

            <div className="form-group">
              <label htmlFor="date">Flight Date</label>
              <input
                type="date"
                id="date"
                name="date"
                value={formData.date}
                onChange={handleInputChange}
                disabled={loading}
              />
            </div>

            <div className="form-group">
              <label htmlFor="aircraft">Aircraft Type</label>
              <select
                id="aircraft"
                name="aircraft"
                value={formData.aircraft}
                onChange={handleInputChange}
                disabled={loading}
              >
                <option value="">Select aircraft type</option>
                {aircraftTypes.map(type => (
                  <option key={type} value={type}>{type}</option>
                ))}
              </select>
            </div>

            <button 
              type="submit" 
              className="generate-btn"
              disabled={loading}
            >
              {loading ? 'Generating...' : 'Generate Vouchers'}
            </button>
          </form>

          {error && (
            <div className="message error-message">
              <div className="message-icon">⚠️</div>
              <div className="message-text">{error}</div>
            </div>
          )}

          {success && generatedSeats && Array.isArray(generatedSeats) && generatedSeats.length > 0 && (
            <div className="message success-message">
              <div className="message-icon">✅</div>
              <div className="message-content">
                <div className="message-text">Vouchers generated successfully!</div>
                <div className="seats-container">
                  <h3>Assigned Seats:</h3>
                  <div className="seats-list">
                    {(generatedSeats || []).map((seat, index) => (
                      <span key={index} className="seat-badge">{seat}</span>
                    ))}
                  </div>
                </div>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default App;