import {service} from "@/api/index";

export function uploadImageApi(file: File, imageType: string) {
    const formData = new FormData();
    formData.set('image', file);
    formData.set('imageType', imageType);
    return service.post('/api/file/image', formData, {
        headers: {
            'Content-Type':'multipart/form-data'
        }
    });
}