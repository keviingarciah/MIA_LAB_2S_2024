import { NavLink, useLocation, useNavigate } from "react-router-dom";
import { useAuthStore } from "../store/authStore";

function Navbar() {
  const location = useLocation();
  const navigate = useNavigate();
  const isLoggedIn = useAuthStore((state) => state.isLoggedIn);
  const logout = useAuthStore((state) => state.logout);

  const handleLogout = () => {
    logout();
    alert("¡Usuario deslogueado exitosamente!"); // Mostrar alerta
    navigate("/login");
  };

  return (
    <div className="bg-green-600">
      <nav className="flex items-center justify-between mx-5 p-4">
        <div className="text-white font-semibold text-4xl flex items-center">
          <p className="font-bold">MIA</p>
        </div>

        <div className="flex">
          {location.pathname !== "/" && (
            <NavLink
              to="/"
              className="ml-auto text-white text-xl font-semibold hover:bg-green-700 px-4 py-4 rounded transition duration-300 ease-in-out"
            >
              Terminal
            </NavLink>
          )}
          {isLoggedIn && location.pathname == "/" && (
            <NavLink
              to="/disks"
              className="ml-auto text-white text-xl font-semibold hover:bg-green-700 px-4 py-4 rounded transition duration-300 ease-in-out"
            >
              Explorador de Archivos
            </NavLink>
          )}
          {!isLoggedIn && location.pathname !== "/login" && (
            <NavLink
              to="/login"
              className="ml-auto text-white text-xl font-semibold hover:bg-green-700 px-4 py-4 rounded transition duration-300 ease-in-out"
            >
              Login
            </NavLink>
          )}
          {isLoggedIn && (
            <button
              onClick={handleLogout}
              className="ml-auto text-white text-xl font-semibold hover:bg-green-700 px-4 py-4 rounded transition duration-300 ease-in-out"
            >
              Logout
            </button>
          )}
        </div>
      </nav>
    </div>
  );
}

export default Navbar;
