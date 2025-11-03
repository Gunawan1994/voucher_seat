import React, { useState } from "react";
import axios from "axios";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";

function App() {
  const [form, setForm] = useState({
    crewName: "",
    crewId: "",
    flightNumber: "",
    flightDate: new Date(),
    aircraftType: "ATR",
  });
  const [seats, setSeats] = useState([]);
  const [message, setMessage] = useState("");
  const [loading, setLoading] = useState(false);

  const API_URL = process.env.REACT_APP_API_URL || "http://localhost:8080";

  const handleChange = (e) => {
    const { name, value } = e.target;
    setForm({ ...form, [name]: value });
  };

  const handleDateChange = (date) => {
    setForm({ ...form, flightDate: date });
  };

  const handleGenerate = async () => {
    setLoading(true);
    setMessage("");
    setSeats([]);

    try {
      // Format date as yyyy-mm-dd
      const flightDateStr = form.flightDate.toISOString().split("T")[0];

      // Check if vouchers exist
      const checkRes = await axios.post(`${API_URL}/api/check`, {
        flightNumber: form.flightNumber,
        date: flightDateStr,
      });

      const exists = checkRes.data?.data?.exists;

      if (exists) {
        setMessage("⚠️ Vouchers already generated for this flight and date.");
        setLoading(false);
        return;
      }

      // Generate vouchers
      const generateRes = await axios.post(`${API_URL}/api/generate`, {
        name: form.crewName,
        id: form.crewId,
        flightNumber: form.flightNumber,
        date: flightDateStr,
        aircraft: form.aircraftType,
      });

      console.log("Generate response:", generateRes.data);

      if (generateRes.data.success) {
        const seats = generateRes.data.data?.seats || [];
        setSeats(seats);
        setMessage(`Vouchers generated successfully! ${seats.length} seats assigned.`);
      } else {
        setMessage(generateRes.data.message || "Failed to generate vouchers.");
      }
    } catch (error) {
      console.error(error);
      setMessage("Error connecting to the server.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={styles.container}>
      <h1>✈️ Voucher Seat Assignment</h1>

      <div style={styles.form}>
        <input
          name="crewName"
          placeholder="Crew Name"
          value={form.crewName}
          onChange={handleChange}
          style={styles.input}
        />
        <input
          name="crewId"
          placeholder="Crew ID"
          value={form.crewId}
          onChange={handleChange}
          style={styles.input}
        />
        <input
          name="flightNumber"
          placeholder="Flight Number"
          value={form.flightNumber}
          onChange={handleChange}
          style={styles.input}
        />

        <DatePicker
          selected={form.flightDate}
          onChange={handleDateChange}
          showTimeSelect
          dateFormat="yyyy-MM-dd HH:mm"
          placeholderText="Select Flight Date & Time"
          style={styles.input}
        />

        <select
          name="aircraftType"
          value={form.aircraftType}
          onChange={handleChange}
          style={styles.input}
        >
          <option value="ATR">ATR</option>
          <option value="Airbus 320">Airbus 320</option>
          <option value="Boeing 737 Max">Boeing 737 Max</option>
        </select>

        <button onClick={handleGenerate} style={styles.button} disabled={loading}>
          {loading ? "Processing..." : "Generate Vouchers"}
        </button>
      </div>

      {message && <p style={styles.message}>{message}</p>}

      {seats.length > 0 && (
        <div style={styles.card}>
          <h3>Generated Seats:</h3>
          <ul>
            {seats.map((seat, i) => (
              <li key={i}>{seat}</li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
}

const styles = {
  container: { fontFamily: "Arial, sans-serif", textAlign: "center", marginTop: "50px" },
  form: {
    display: "flex",
    flexDirection: "column",
    gap: "10px",
    alignItems: "center",
    maxWidth: "400px",
    margin: "0 auto",
  },
  input: { padding: "10px", width: "100%", borderRadius: "5px", border: "1px solid #ccc" },
  button: {
    backgroundColor: "#007bff",
    color: "white",
    border: "none",
    padding: "10px 20px",
    cursor: "pointer",
    fontSize: "16px",
    borderRadius: "5px",
  },
  message: { marginTop: "20px", fontWeight: "bold" },
  card: {
    marginTop: "20px",
    backgroundColor: "#f8f9fa",
    padding: "20px",
    borderRadius: "10px",
    maxWidth: "300px",
    margin: "20px auto",
    boxShadow: "0px 2px 5px rgba(0,0,0,0.1)",
  },
};

export default App;
