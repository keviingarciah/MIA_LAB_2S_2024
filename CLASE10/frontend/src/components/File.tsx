import fileIcon from "../assets/file-icon.png";

interface FileProps {
  name: string;
}

const File: React.FC<FileProps> = ({ name }) => {
  return (
    <div className="p-4 bg-gray-200 rounded-lg shadow-md w-28 text-center">
      <img src={fileIcon} alt="Folder Icon" className="mx-auto mb-2 w-10" />
      <div className="text-blue-500 text-base">{name}</div>
    </div>
  );
};

export default File;
