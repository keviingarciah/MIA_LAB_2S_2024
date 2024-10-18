import folderIcon from "../assets/folder-icon.png";

interface FolderProps {
  name: string;
}

const Folder: React.FC<FolderProps> = ({ name }) => {
  return (
    <div className="p-4 bg-gray-200 rounded-lg shadow-md w-28 text-center">
      <img src={folderIcon} alt="Folder Icon" className="mx-auto mb-2 w-10" />
      <div className="text-yellow-500 text-base">{name}</div>
    </div>
  );
};

export default Folder;
