import { useParams, Link } from 'react-router-dom';
import { useState, useEffect } from 'react';

function ArticlePage() {
  const [article, setArticle] = useState(null);
  const [notFound, setNotFound] = useState(false);
  const [loading, setLoading] = useState(true);
  const { name } = useParams();

  useEffect(() => {
    fetch(`/articles/${name}`)
      .then(response => {
        if (!response.ok) {
          setNotFound(true);
          return {};
        }
        return response.json();
      })
      .then(data => {
        setArticle(data);
        setLoading(false);
      });
  }, [name]);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (notFound) {
    return (
      <div>
        <h1>{name}</h1>
        <p>No article with this exact name found. Use Edit button in the header to add it.</p>
        <Link to={`/edit/${name}`}>Edit</Link>
      </div>
    );
  }

  return (
    <div>
      <h1>{article.name}</h1>
      <Link to={`/edit/${name}`}>Edit</Link>
      
      <div dangerouslySetInnerHTML={{ __html: article.content }} />
      <Link to="/">All Articles</Link>
    </div>
  );
}

export default ArticlePage;