import React, {useState, useEffect} from 'react';

import 'bootstrap/dist/css/bootstrap.css'

import {Button, Card, Row, Col} from 'react-bootstrap'

const Entry =({entryData, setChangeIngredient, deleteSingleEntry, setChangeEntry}) => {
    return(
        <Card className='p-2'>
                <Col>Dish:{entryData !== undefined && entryData.dish}</Col>
                <Col>Ingredients:{entryData !== undefined && entryData.ingredients}</Col>
                <Col>Calories:{entryData !== undefined && entryData.calories}</Col>
                <Col>Fat:{entryData !== undefined && entryData.fat}</Col>

                <div className='flex flex-row'>
                <Col><Button onClick={()=> deleteSingleEntry(entryData._id)}>delete entry</Button></Col>
                {/* <Col><Button onClick={()=> changeIngredient()}>change ingredients</Button></Col> */}
                <Col><Button onClick={()=> changeEntry()}>change entry</Button></Col>
                    
                </div>

        </Card>
    )

    function changeIngredient(){
        setChangeIngredient(
            {
                "change": true,
                "id":entryData._id
            }
        )
    }

    function changeEntry(){
        setChangeEntry(
            {
                "change": true,
                "id":entryData._id
            }
        )
    }
}

export default Entry