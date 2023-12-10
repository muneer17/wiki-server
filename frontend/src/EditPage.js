import { useParams, Link, useNavigate } from 'react-router-dom';
import { useState, useEffect } from 'react';

function EditPage() {
  const [article, setArticle] = useState(null);
  const [content, setContent] = useState('');
  const { name } = useParams();
  const navigate = useNavigate();

  useEffect(() => {
    fetch(`/articles/${name}`)
      .then(response => {
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        return response.text();
      })
      .then(data => {
        if (data) {
          const jsonData = JSON.parse(data);
          setArticle(jsonData);
          setContent(jsonData.content);
        }
      })
      .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
      });
  }, [name]);

  function handleSave() {
    fetch(`/articles/${name}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ content }),
    })
    .then(response => {
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      return response.text();
    })
    .then(data => {
      if (data) {
        try {
          return JSON.parse(data);
        } catch (err) {
          console.error('The server did not return valid JSON:', data);
          throw err;
        }
      }
    })
    .then(() => {
      navigate(`/${name}`);
    })
    .catch(error => {
      console.error('There was a problem with the fetch operation:', error);
    });
  }

  function handleCancel() {
    navigate(`/${name}`);
  }

return (
  <div>
    <h1>{name}</h1>
    <Link to={`/${name}`}>View</Link>
    <br />
    <textarea name="content" value={content} onChange={e => setContent(e.target.value)} />
    <br />
    <button onClick={handleSave}>Save</button>
    <br />
    <button onClick={handleCancel}>Cancel</button>
  </div>
);

}

export default EditPage;