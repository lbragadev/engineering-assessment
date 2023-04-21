"""
This script fetches data from sfgov SODA API. It then processes the data and inserts the 
data to the food_trucks table.
"""

import psycopg2
import requests
from datetime import datetime


def get_food_trucks():
    # Fetches list of food trucks from sfgov api and returns as json list.

    url = 'https://data.sfgov.org/resource/rqzj-sfat.json'
    response = requests.get(url)    
    data = response.json()
    return data

def connect_to_db():
    # Connects to postgres database and returns connection.
    conn = psycopg2.connect(host="pgdb", port="5432", database="elc_db", user="postgres", password="postgres")
    return conn

def get_location_ids():
    # Fetches all location_ids from food_trucks table and returns as a list.
    conn = connect_to_db()
    cur = conn.cursor()
    cur.execute("SELECT location_id FROM food_trucks")
    result = []
    for row in cur.fetchall():
        location_id = row[0]
        result.append(str(location_id))
    conn.close()
    return result

def process_food_trucks():
    """
    This function fetches all food trucks from sfgov api. It then fetches all location_ids of existing food_trucks in the food_trucks table.
    Each food_truck is processed and inserted into the food_trucks table. The location_id of each food_truck is checked so that an already
    existing food_truck in the table does not get re-inserted. This allows the script to be re-runnable in order to fetch for new food_trucks 
    without duplicating old ones.
    """


    conn = connect_to_db()
    cur = conn.cursor()
    food_trucks = get_food_trucks()

    #TODO  
        #ON CONFLICT UPDATE instead of skip

    #get list of all location ids from db
    location_ids = get_location_ids()

    for food_truck in food_trucks:
        dt = datetime.now()

        # skip food_trucks that already exists in the db.
        if food_truck['objectid'] in location_ids:
            continue
        status = ""
        facilitytype = ""
        locationdescription = ""
        fooditems = ""
        latitude = ""
        longitude = ""
        if "status" in food_truck:
            status = food_truck['status']
        if "facilitytype" in food_truck:
            facilitytype = food_truck['facilitytype']
        if "locationdescription" in food_truck:
            locationdescription = food_truck['locationdescription']
        if "fooditems" in food_truck:
            fooditems = food_truck['fooditems']
        if "latitude" in food_truck:
            latitude = food_truck['latitude']
        if "longitude" in food_truck:
            longitude = food_truck['longitude']

        insert_query = '''
            INSERT INTO food_trucks (location_id, name, address, status, facility_type, 
            location_description, food_items, latitude, longitude, created_at) 
            VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
            '''
        cur.execute(insert_query, (food_truck['objectid'], food_truck['applicant'], food_truck['address'], status,
                    facilitytype, locationdescription, fooditems, latitude, longitude, dt))
    conn.commit()
    conn.close()

if __name__ == '__main__':
    process_food_trucks()
