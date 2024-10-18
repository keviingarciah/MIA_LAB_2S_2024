import { useParams } from "react-router-dom";
import Partition from "../components/Partition";

const partitionsData = [
  { id: "1", name: "Partición 1", fileSystem: "NTFS" },
  { id: "2", name: "Partición 2", fileSystem: "EXT4" },
  { id: "3", name: "Partición 3", fileSystem: "FAT32" },
];

const PartitionsPage: React.FC = () => {
  const { diskId } = useParams<{ diskId: string }>();

  return (
    <div className="flex-grow flex items-center justify-center p-44">
      <div className="w-full max-w-3xl p-8 bg-white rounded-lg shadow-md">
        <h2 className="text-2xl font-bold mb-4 text-gray-800">
          Particiones del Disco {diskId}
        </h2>
        <div className="flex flex-wrap gap-4">
          {partitionsData.map((partition) => (
            <Partition
              key={partition.id}
              id={partition.id}
              name={partition.name}
            />
          ))}
        </div>
      </div>
    </div>
  );
};

export default PartitionsPage;
