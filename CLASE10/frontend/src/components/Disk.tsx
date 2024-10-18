import { useNavigate } from "react-router-dom";
import diskIcon from "../assets/disk-icon.png";

interface DiskProps {
  id: string;
  name: string;
}

const Disk: React.FC<DiskProps> = ({ id, name }) => {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate(`/partitions/${id}`);
  };

  return (
    <div
      onClick={handleClick}
      className="cursor-pointer p-4 bg-gray-200 rounded-lg shadow-md hover:bg-gray-300 transition duration-300 ease-in-out"
    >
      <img src={diskIcon} alt="Folder Icon" className="mx-auto mb-2 w-16" />
      <h3 className="text-base font-medium">{name}</h3>
    </div>
  );
};

export default Disk;
