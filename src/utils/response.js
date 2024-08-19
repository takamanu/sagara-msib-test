const response = (statusCode, data, message, res) => {
  res.status(statusCode).json({
    message,
    payload: data,
  });
}

module.exports = response;