import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuthStore } from "../store/authStore";

const Login: React.FC = () => {
  const [partitionId, setPartitionId] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const navigate = useNavigate();

  const login = useAuthStore((state) => state.login);

  const handlePartitionIdChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPartitionId(e.target.value);
  };

  const handleUsernameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUsername(e.target.value);
  };

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    // Aquí puedes manejar el envío del formulario, por ejemplo, hacer una solicitud a una API
    console.log("Partition ID:", partitionId);
    console.log("Username:", username);
    console.log("Password:", password);
    alert("¡Usuario logueado exitosamente!"); // Mostrar alerta
    handleLogin();
  };

  const handleLogin = () => {
    // Aquí puedes agregar la lógica de autenticación
    login();
    console.log("User is logged in:", useAuthStore.getState().isLoggedIn);

    navigate("/"); // Redirecciona a la página principal después del login
  };

  return (
    <div className="flex justify-center items-center h-full p-32">
      <form
        onSubmit={handleSubmit}
        className="bg-white p-6 rounded shadow-md w-80"
      >
        <h2 className="text-2xl mb-4">Login</h2>
        <div className="mb-4">
          <label className="block text-gray-700">ID Partición</label>
          <input
            type="text"
            value={partitionId}
            onChange={handlePartitionIdChange}
            className="w-full px-3 py-2 border rounded"
            required
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Usuario</label>
          <input
            type="text"
            value={username}
            onChange={handleUsernameChange}
            className="w-full px-3 py-2 border rounded"
            required
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Contraseña</label>
          <input
            type="password"
            value={password}
            onChange={handlePasswordChange}
            className="w-full px-3 py-2 border rounded"
            required
          />
        </div>
        <button
          type="submit"
          className="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-700"
        >
          Iniciar Sesión
        </button>
      </form>
    </div>
  );
};

export default Login;
