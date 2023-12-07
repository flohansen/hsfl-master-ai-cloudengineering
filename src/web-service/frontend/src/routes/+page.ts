import {handleErrors} from "../assets/helper/handleErrors";

// Loads all shopping lists of the current user.
export const load = (): Promise<object> => {
    const id: number = 2; // TODO: dynamic user id of the current logged-in user
    const apiUrl: string = `/api/v1/shoppinglist/${id}`;

    return fetch(apiUrl)
        .then(handleErrors)
        .then(lists => {
            return {
                lists: lists.slice(0, 3),
                metaTitle: 'Startseite',
                headline: 'Price Whisper',
            };
        })
        .catch(error => {
            console.error("Failed to fetch shopping lists data:", error.message);
            return {
                metaTitle: 'Error',
                headline: 'Price Whisper',
            };
        });
};
