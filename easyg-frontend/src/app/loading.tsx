"use client"
import Spinner from 'react-bootstrap/Spinner';

function Loading() {
  return (
    <div style={
      {
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "20vh",
      }
    }>
      <Spinner animation="border" role="status">
        <span className="visually-hidden">Loading...</span>
      </Spinner>
    </div>
  );
}

export default Loading;