import JoditEditor, { Jodit } from 'jodit-react';
import React, { useState, useRef, useMemo } from 'react';

const JoditEditorComponent: React.FC = () => {
  const editor = useRef<Jodit | null>(null);
  const [content, setContent] = useState<string>('');

  const config = useMemo(
    () => ({
      readonly: false,
      placeholder: 'Start typing...',
      uploader: {
        insertImageAsBase64URI: true,
      },
    }),
    [],
  );

  const handleContentChange = (newContent: string) => {
    setContent(newContent);
  };

  return (
    <JoditEditor
      ref={editor}
      value={content}
      config={config}
      onBlur={handleContentChange}
      onChange={handleContentChange}
    />
  );
};

export default JoditEditorComponent;
