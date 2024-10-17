import React, { useState, useRef } from "react";

export default function Terminal() {
  const [inputText, setInputText] = useState("");
  const [outputText, setOutputText] = useState("");
  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = (e) => {
        const content = e.target?.result as string;
        setInputText(content);
      };
      reader.readAsText(file);
    }
  };

  const handleExecute = async () => {
    try {
      const response = await fetch("http://localhost:3000/analyze", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ command: inputText }),
      });

      if (!response.ok) {
        throw new Error("Network response was not ok");
      }

      const data = await response.json();
      const results = data.results.join("\n");
      setOutputText(results);
    } catch (error) {
      console.error("Error:", error);
      setOutputText(`Error: ${error}`);
    }
  };

  return (
    <div className="flex-grow flex items-center justify-center p-4">
      <div className="w-full max-w-3xl p-8 bg-white rounded-lg shadow-md">
        <div className="mb-4">
          <textarea
            className="w-full h-48 p-2 border border-gray-300 rounded-md resize-none focus:outline-none focus:ring-2 focus:ring-blue-500"
            value={inputText}
            onChange={(e) => setInputText(e.target.value)}
            placeholder="Terminal de entrada"
          />
        </div>
        <div className="mb-4">
          <textarea
            className="w-full h-48 p-2 border border-gray-300 rounded-md resize-none bg-gray-100 focus:outline-none"
            value={outputText}
            readOnly
            placeholder="Terminal de salida"
          />
        </div>
        <div className="flex justify-between">
          <div>
            <input
              type="file"
              ref={fileInputRef}
              onChange={handleFileChange}
              className="hidden"
              accept=".txt"
            />
            <button
              onClick={() => fileInputRef.current?.click()}
              className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              Examinar
            </button>
          </div>
          <button
            onClick={handleExecute}
            className="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-600"
          >
            Ejecutar
          </button>
        </div>
      </div>
    </div>
  );
}
