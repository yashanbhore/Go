import React, { useState, useEffect } from "react";
import axios from "axios";
import {
  Button,
  Form,
  Container,
  Modal,
  Row,
  Col,
  Card,
} from "react-bootstrap";
import { Doughnut } from "react-chartjs-2";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import Entry from "./single.entry.component";
import "../App.css"; // Import custom CSS for additional styling

// Register necessary elements for Chart.js
ChartJS.register(ArcElement, Tooltip, Legend);

const Entries = () => {
  const [entries, setEntries] = useState([]);
  const [refreshData, setRefreshData] = useState(false);
  const [changeEntry, setChangeEntry] = useState({ change: false, id: 0 });
  const [changeIngredient, setChangeIngredient] = useState({
    change: false,
    id: 0,
  });
  const [newIngredientName, setNewIngredientName] = useState("");
  const [addNewEntry, setAddNewEntry] = useState(false);
  const [newEntry, setNewEntry] = useState({
    dish: "",
    ingredients: "",
    calories: 0,
    fat: 0,
    protein: 0,
    carbohydrates: 0,
  });

  useEffect(() => {
    getAllEntries();
  }, []);

  useEffect(() => {
    if (refreshData) {
      setRefreshData(false);
      getAllEntries();
    }
  }, [refreshData]);

  const handleChange = (e) => {
    setNewEntry({ ...newEntry, [e.target.name]: e.target.value });
  };

  function changeIngredientForEntry(){
    changeIngredient.change = false
    var url = "http://localhost:8000/ingredient/update/" + changeIngredient.id
    axios.put(url, {
        "ingredients": newIngredientName
    }).then(response => {
        console.log(response.status)
        if(response.status == 200 ){
            setRefreshData(true)
        }
    })
}

function changeSingleEntry(){
    changeEntry.change = false;
    var url = "http://localhost:8000/entry/update/" + changeEntry.id
    axios.put(url, newEntry)
    .then(response =>{
        if(response.status == 200){
            setRefreshData(true)
        }
    })
}

function addSingleEntry(){
    setAddNewEntry(false)
    var url = "http://localhost:8000/entry/create"
    axios.post(url, {
        "ingredients":newEntry.ingredients,
        "dish": newEntry.dish,
        "calories": newEntry.calories,
        "fat": parseFloat(newEntry.fat)
    }).then(response => {
        if(response.status == 200){
            setRefreshData(true)
        }
    })
}

function deleteSingleEntry(id){
    var url = "http://localhost:8000/entry/delete/" + id
    axios.delete(url, {

    }).then(response => {
        if (response.status == 200){
            setRefreshData(true)
        }
    })
}

function getAllEntries(){
    var url = "http://localhost:8000/entries"
    axios.get(url, {
        reponseType: 'json'
    }).then(response => {
        if(response.status == 200){
            setEntries(response.data)
        }
    })
}
  const calculateTotal = (property) => {
    return entries.reduce((total, entry) => total + parseFloat(entry[property]), 0);
  };

  const data = {
    labels: ["Calories", "Protein", "Fats", "Carbohydrates"],
    datasets: [
      {
        data: [
          calculateTotal("calories"),
          calculateTotal("protein"),
          calculateTotal("fat"),
          calculateTotal("carbohydrates"),
        ],
        backgroundColor: ["#FF6384", "#36A2EB", "#FFCE56", "#4BC0C0"],
        hoverBackgroundColor: ["#FF6384", "#36A2EB", "#FFCE56", "#4BC0C0"],
      },
    ],
  };

  return (
    <div className="py-5" style={{ backgroundColor: "#f5f5f5" }}>
      <Container className="p-4 mb-2 flex flex-col justify-center">
        <Button variant="primary" onClick={() => setAddNewEntry(true)}>
          Track Today's Calories
        </Button>
        <div className="flex">
          <Container className="p-4 mb-2 flex justify-center small-chart-container">
            <Doughnut data={data} />
          </Container>
        </div>
      </Container>
      <Container>
        <Row className="p-2">
          {entries &&
            entries.map((entry, i) => (
              <Col key={i} sm={12} md={6} lg={4} className="mb-4">
                    <Entry
                      entryData={entry}
                      deleteSingleEntry={deleteSingleEntry}
                      setChangeIngredient={setChangeIngredient}
                      setChangeEntry={setChangeEntry}
                    />
              </Col>
            ))}
        </Row>
      </Container>

      <Modal show={addNewEntry} onHide={() => setAddNewEntry(false)} centered>
        <Modal.Header closeButton>
          <Modal.Title>Add Calorie Entry</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
            <Form.Group>
              <Form.Label>Dish</Form.Label>
              <Form.Control name="dish" onChange={handleChange} />
              <Form.Label>Ingredients</Form.Label>
              <Form.Control name="ingredients" onChange={handleChange} />
              <Form.Label>Calories</Form.Label>
              <Form.Control name="calories" onChange={handleChange} />
              <Form.Label>Fat</Form.Label>
              <Form.Control name="fat" onChange={handleChange} />
              <Form.Label>Protein</Form.Label>
              <Form.Control name="protein" onChange={handleChange} />
              <Form.Label>Carbohydrates</Form.Label>
              <Form.Control name="carbohydrates" onChange={handleChange} />
            </Form.Group>
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="success" onClick={addSingleEntry}>
            Add
          </Button>
          <Button variant="secondary" onClick={() => setAddNewEntry(false)}>
            Cancel
          </Button>
        </Modal.Footer>
      </Modal>
      <Modal
        show={changeIngredient.change}
        onHide={() => setChangeIngredient({ change: false, id: 0 })}
        centered
      >
        <Modal.Header closeButton>
          <Modal.Title>Change Ingredients</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group>
            <Form.Label>New Ingredients</Form.Label>
            <Form.Control
              onChange={(event) => setNewIngredientName(event.target.value)}
            />
          </Form.Group>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="success" onClick={changeIngredientForEntry}>
            Change
          </Button>
          <Button
            variant="secondary"
            onClick={() => setChangeIngredient({ change: false, id: 0 })}
          >
            Cancel
          </Button>
        </Modal.Footer>
      </Modal>
      <Modal
        show={changeEntry.change}
        onHide={() => setChangeEntry({ change: false, id: 0 })}
        centered
      >
        <Modal.Header closeButton>
          <Modal.Title>Change Entry</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group>
            <Form.Label>Dish</Form.Label>
            <Form.Control name="dish" onChange={handleChange} />
            <Form.Label>Ingredients</Form.Label>
            <Form.Control name="ingredients" onChange={handleChange} />
            <Form.Label>Calories</Form.Label>
            <Form.Control name="calories" onChange={handleChange} />
            <Form.Label>Fat</Form.Label>
            <Form.Control name="fat" onChange={handleChange} />
            <Form.Label>Protein</Form.Label>
            <Form.Control name="protein" onChange={handleChange} />
            <Form.Label>Carbohydrates</Form.Label>
            <Form.Control name="carbohydrates" onChange={handleChange} />
          </Form.Group>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="success" onClick={changeSingleEntry}>
            Change
          </Button>
          <Button
            variant="secondary"
            onClick={() => setChangeEntry({ change: false, id: 0 })}
          >
            Cancel
          </Button>
        </Modal.Footer>
      </Modal>
    </div>
  );
};

export default Entries;
