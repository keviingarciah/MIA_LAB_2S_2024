import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import Folder from "../components/Folder";
import File from "../components/File";

const fileSystemsData: { [key: string]: string } = {
  "1": "NTFS",
  "2": "EXT4",
  "3": "FAT32",
};

const FileSystemPage: React.FC = () => {
  const { partitionId } = useParams<{ partitionId: string }>();
  const [path, setPath] = useState("/");
  const [results, setResults] = useState<{ type: string; name: string }[]>([]);

  const fileSystem =
    partitionId && fileSystemsData[partitionId]
      ? fileSystemsData[partitionId]
      : "Desconocido";

  const handleSearch = () => {
    // Simular resultados de búsqueda
    const simulatedResults = [
      { type: "folder", name: "Carpeta 1" },
      { type: "folder", name: "Carpeta 2" },
      { type: "file", name: "Texto 1.txt" },
      { type: "file", name: "Texto 2.txt" },
    ];
    setResults(simulatedResults);
  };

  useEffect(() => {
    handleSearch();
  }, []);

  return (
    <div className="flex-grow flex flex-col items-center justify-center p-16">
      <div className="w-full max-w-3xl p-8 bg-white rounded-lg shadow-md">
        <h2 className="text-2xl font-bold mb-4 text-gray-800">
          Sistema de Archivos de la Partición {partitionId}
        </h2>
        <p className="text-gray-700 mb-4">Sistema de Archivos: {fileSystem}</p>
        <div className="flex mb-4">
          <input
            type="text"
            value={path}
            onChange={(e) => setPath(e.target.value)}
            placeholder="Ingrese el path"
            className="flex-grow p-2 border border-gray-300 rounded-l-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <button
            onClick={handleSearch}
            className="p-2 bg-blue-500 text-white rounded-r-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            Buscar
          </button>
        </div>
        <div className="flex flex-wrap gap-4">
          {results.map((item, index) =>
            item.type === "folder" ? (
              <Folder key={index} name={item.name} />
            ) : (
              <File key={index} name={item.name} />
            )
          )}
        </div>
      </div>
    </div>
  );
};

export default FileSystemPage;
