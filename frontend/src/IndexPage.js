import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';

function IndexPage() {
  const [articles, setArticles] = useState([]);
  
  useEffect(() => {
    fetch('/articles/')
      .then(response => {
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        return response.text();
      })
      .then(data => {
        try {
          return JSON.parse(data);
        } catch (err) {
          console.error('The server did not return valid JSON:', data);
          throw err;
        }
      })
      .then(articles => setArticles(articles))
      .catch(error => console.error('There was a problem with the fetch operation:', error));
  }, []);

  console.log(articles)
  
  return (
    <div>
      <h1>Articles</h1>
      <ul>
        {articles.map(article => (
          <li key={article.name}>
            <Link to={`/${article.name}`}>{article.name}</Link>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default IndexPage;