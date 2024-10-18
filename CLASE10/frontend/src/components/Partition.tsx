import { useNavigate } from "react-router-dom";
import partitionIcon from "../assets/partition-icon.png";

interface PartitionProps {
  id: string;
  name: string;
}

const Partition: React.FC<PartitionProps> = ({ id, name }) => {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate(`/filesystem/${id}`);
  };

  return (
    <div
      onClick={handleClick}
      className="cursor-pointer p-4 bg-gray-200 rounded-lg shadow-md hover:bg-gray-300 transition duration-300 ease-in-out"
    >
      <img
        src={partitionIcon}
        alt="Folder Icon"
        className="mx-auto mb-2 w-16"
      />
      <h3 className="text-xl font-medium">{name}</h3>
    </div>
  );
};

export default Partition;
