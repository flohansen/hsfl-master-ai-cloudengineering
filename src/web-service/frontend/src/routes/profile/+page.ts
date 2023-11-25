import {handleErrors} from "../../assets/helper/handleErrors";

export const load = async (): Promise<object> => {
    const userId: number = 2; // TODO: add real current user id
    const apiUrl: string = `/api/v1/user/${userId}`;

    return fetch(apiUrl)
        .then(handleErrors)
        .then(user => {
            return {
                user: user ?? [],
                metaTitle: 'Deine Profil-Einstellungen',
                headline: 'Dein Profil',
            };
        })
        .catch(error => {
            console.error("Failed to fetch user data:", error.message);
            return {
                user: [],
                metaTitle: 'Deine Profil-Einstellungen',
                headline: 'Dein Profil',
            };
        });
};